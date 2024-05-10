package database_postgres

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	entities_game_v1 "github.com/teyz/songify-svc/internal/entities/game/v1"
	"github.com/teyz/songify-svc/internal/pkg/errors"
)

func (d *dbClient) CreateGame(ctx context.Context, songID string) (*entities_game_v1.Game, error) {
	game := &entities_game_v1.Game{}

	id := GenerateDataPrefixWithULID(Game)

	err := d.connection.DB.QueryRowContext(ctx,
		`
			INSERT INTO games (id, song_id)
			VALUES ($1, $2)
			RETURNING id, song_id, is_active, created_at
		`,
		id, songID).Scan(
		&game.ID,
		&game.SongID,
		&game.IsActive,
		&game.CreatedAt,
	)
	if err != nil {
		log.Error().Err(err).
			Str("id", id).
			Msgf("database.postgres.dbClient.CreateGame: failed to create game: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.CreateGame: failed to create game: %v", err.Error()))
	}

	return game, nil
}

func (d *dbClient) GetCurrentGame(ctx context.Context) (*entities_game_v1.Game, error) {
	game := &entities_game_v1.Game{}

	err := d.connection.DB.QueryRowContext(ctx,
		`
			SELECT id, song_id, is_active, created_at
			FROM games
			WHERE is_active = true
			ORDER BY created_at DESC
		`).Scan(
		&game.ID,
		&game.SongID,
		&game.IsActive,
		&game.CreatedAt,
	)
	if err != nil {
		log.Error().Err(err).
			Msgf("database.postgres.dbClient.GetCurrentGame: failed to get current game: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.GetCurrentGame: failed to get current game: %v", err.Error()))
	}

	return game, nil
}