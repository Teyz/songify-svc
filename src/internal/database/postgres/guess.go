package database_postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"

	entities_guess_v1 "github.com/teyz/songify-svc/internal/entities/guess/v1"
	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
	"github.com/teyz/songify-svc/pkg/errors"
)

func (d *dbClient) CheckIfUserCanGuess(ctx context.Context, userID string, gameID string) (int32, error) {
	var countUserGuesses int32

	err := d.connection.DB.QueryRowContext(ctx, `
		SELECT
			COUNT(*)
		FROM
			guesses
		WHERE
			user_id = $1
			AND game_id = $2
	`, userID, gameID).Scan(&countUserGuesses)
	if err != nil {
		log.Error().Err(err).
			Str("user_id", userID).
			Str("game_id", gameID).
			Msgf("database.postgres.dbClient.CheckIfUserCanGuess: failed to check if user can guess: %v", err.Error())
		return 0, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.CheckIfUserCanGuess: failed to check if user can guess: %v", err.Error()))
	}

	return countUserGuesses, nil
}

func (d *dbClient) CheckGuess(ctx context.Context, songID string) (*entities_song_v1.Song_Guess, error) {
	song := &entities_song_v1.Song_Guess{}

	err := d.connection.DB.QueryRowContext(ctx, `
		SELECT 
			id, 
			title, 
			artist
		FROM
			songs
		WHERE
			id = $1
	`, songID).Scan(
		&song.ID,
		&song.Title,
		&song.Artist,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).
				Str("song_id", songID).
				Msgf("database.postgres.dbClient.CheckGuess: song not found")
			return nil, errors.NewNotFoundError("database.postgres.dbClient.CheckGuess: song not found")
		}

		log.Error().Err(err).
			Str("song_id", songID).
			Msgf("database.postgres.dbClient.CheckGuess: failed to check guess: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.CheckGuess: failed to check guess: %v", err.Error()))
	}

	return song, nil
}

func (d *dbClient) AddGuess(ctx context.Context, userID string, guess *entities_guess_v1.Guess) error {
	id := GenerateDataPrefixWithULID(Guess)

	_, err := d.connection.DB.ExecContext(ctx, `
		INSERT INTO guesses (id, user_id, game_id, title, is_title_correct, artist, is_artist_correct)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, id, userID, guess.GameID, guess.Title, guess.IsTitleCorrect, guess.Artist, guess.IsArtistCorrect)
	if err != nil {
		log.Error().Err(err).
			Str("game_id", guess.GameID).
			Str("user_id", userID).
			Msgf("database.postgres.dbClient.AddGuess: failed to add guess: %v", err.Error())
		return errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.AddGuess: failed to add guess: %v", err.Error()))
	}

	return nil
}

func (d *dbClient) GetGuessesByUserIDForGame(ctx context.Context, userID string, gameID string) (*entities_guess_v1.Guesses, error) {
	rows, err := d.connection.DB.QueryContext(ctx, `
		SELECT
			title,
			is_title_correct,
			artist,
			is_artist_correct
		FROM
			guesses
		WHERE
			user_id = $1
			AND game_id = $2
	`, userID, gameID)
	if err != nil {
		log.Error().Err(err).
			Str("user_id", userID).
			Str("game_id", gameID).
			Msgf("database.postgres.dbClient.GetGuessesByUserIDForGame: failed to get guesses: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.GetGuessesByUserIDForGame: failed to get guesses: %v", err.Error()))
	}
	defer rows.Close()

	guesses := make([]*entities_guess_v1.Guess, 0)
	isTitleCorrect := false
	isArtistCorrect := false

	for rows.Next() {
		guess := &entities_guess_v1.Guess{}

		err := rows.Scan(
			&guess.Title,
			&guess.IsTitleCorrect,
			&guess.Artist,
			&guess.IsArtistCorrect,
		)
		if err != nil {
			log.Error().Err(err).
				Str("user_id", userID).
				Str("game_id", gameID).
				Msgf("database.postgres.dbClient.GetGuessesByUserIDForGame: failed to scan guess: %v", err.Error())
			return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.GetGuessesByUserIDForGame: failed to scan guess: %v", err.Error()))
		}

		if guess.IsTitleCorrect {
			isTitleCorrect = true
		}

		if guess.IsArtistCorrect {
			isArtistCorrect = true
		}

		guesses = append(guesses, guess)
	}

	return &entities_guess_v1.Guesses{
		IsArtistCorrect: isArtistCorrect,
		IsTitleCorrect:  isTitleCorrect,
		Guesses:         guesses,
	}, nil
}
