package database_postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"
	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
	"github.com/teyz/songify-svc/internal/pkg/errors"
)

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
			return nil, errors.NewNotFoundError(fmt.Sprintf("database.postgres.dbClient.CheckGuess: song not found"))
		}

		log.Error().Err(err).
			Str("song_id", songID).
			Msgf("database.postgres.dbClient.CheckGuess: failed to check guess: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.CheckGuess: failed to check guess: %v", err.Error()))
	}

	return song, nil
}
