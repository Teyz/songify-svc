package database_postgres

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"

	"github.com/rs/zerolog/log"

	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
	"github.com/teyz/songify-svc/pkg/errors"
)

func (d *dbClient) GetSongByID(ctx context.Context, id string) (*entities_song_v1.Song, error) {
	song := &entities_song_v1.Song{}

	err := d.connection.DB.QueryRowContext(ctx,
		`SELECT
			id,
			title,
			artist,
			artist_image_url,
			lyrics,
			image_url,
			released_year,
			musical_style,
			created_at, 
			updated_at
		FROM
			songs
		WHERE
			id = $1
		`,
		id).Scan(
		&song.ID,
		&song.Title,
		&song.Artist,
		&song.ArtistImageURL,
		&song.Lyrics,
		&song.ImageURL,
		&song.ReleasedYear,
		&song.MusicalStyle,
		&song.CreatedAt,
		&song.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).
				Str("id", id).
				Msgf("database.postgres.dbClient.GetSongByID: song with id: %s not found", id)
			return nil, errors.NewNotFoundError(fmt.Sprintf("database.postgres.dbClient.GetSongByID: song with id: %s not found", id))
		}

		log.Error().Err(err).
			Str("id", id).
			Msgf("database.postgres.dbClient.GetSongByID: failed to get song by id: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.GetSongByID: failed to get song by id: %v", err.Error()))
	}

	return song, nil
}

func (d *dbClient) GetRandomSong(ctx context.Context) (*entities_song_v1.Song, error) {
	song := &entities_song_v1.Song{}

	songsLength, err := d.getSongsLength(ctx)
	if err != nil {
		return nil, err
	}

	offset := rand.Intn(songsLength)

	err = d.connection.DB.QueryRowContext(ctx,
		`SELECT
			id,
			title,
			artist,
			artist_image_url,
			lyrics,
			image_url,
			released_year,
			musical_style,
			created_at, 
			updated_at
		FROM
			songs
		WHERE
			has_been_daily_used IS FALSE
		OFFSET 
			$1
		`,
		offset).Scan(
		&song.ID,
		&song.Title,
		&song.Artist,
		&song.ArtistImageURL,
		&song.Lyrics,
		&song.ImageURL,
		&song.ReleasedYear,
		&song.MusicalStyle,
		&song.CreatedAt,
		&song.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).
				Msgf("database.postgres.dbClient.GetRandomSong: song with offset: %d not found", offset)
			return nil, errors.NewNotFoundError(fmt.Sprintf("database.postgres.dbClient.GetRandomSong: song with offset: %d not found", offset))
		}

		log.Error().Err(err).
			Msgf("database.postgres.dbClient.GetRandomSong: failed to get random song: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.GetRandomSong: failed to get random song: %v", err.Error()))
	}

	err = d.updateSongDailyUsage(ctx, song.ID)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (d *dbClient) getSongsLength(ctx context.Context) (int, error) {
	songLenght := 0

	err := d.connection.DB.QueryRowContext(ctx,
		`SELECT
			COUNT(*)
		FROM
			songs
		WHERE
			has_been_daily_used IS FALSE;
		`).Scan(
		&songLenght,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).
				Msgf("database.postgres.dbClient.getSongsLength: no songs in the database")
			return 0, errors.NewNotFoundError("database.postgres.dbClient.getSongsLength: no songs in the database")
		}

		log.Error().Err(err).
			Msgf("database.postgres.dbClient.getSongsLength: failed to get songs length: %v", err.Error())
		return 0, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.getSongsLength: failed to get songs length: %v", err.Error()))
	}

	return songLenght, nil
}

func (d *dbClient) updateSongDailyUsage(ctx context.Context, id string) error {
	_, err := d.connection.DB.ExecContext(ctx,
		`UPDATE
			songs
		SET
			has_been_daily_used = TRUE
		WHERE
			id = $1
		`,
		id)
	if err != nil {
		log.Error().Err(err).
			Str("id", id).
			Msgf("database.postgres.dbClient.UpdateSongDailyUsage: failed to update song daily usage: %v", err.Error())
		return errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.UpdateSongDailyUsage: failed to update song daily usage: %v", err.Error()))
	}

	return nil
}
