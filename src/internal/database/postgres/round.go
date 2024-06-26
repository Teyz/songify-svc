package database_postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"

	entities_round_v1 "github.com/teyz/songify-svc/internal/entities/round/v1"
	"github.com/teyz/songify-svc/pkg/constants"
	"github.com/teyz/songify-svc/pkg/errors"
)

func (d *dbClient) CreateRound(ctx context.Context, userID string, gameID string) (*entities_round_v1.Round, error) {
	round := &entities_round_v1.Round{}

	id := GenerateDataPrefixWithULID(Round)

	err := d.connection.DB.QueryRowContext(ctx, `
		INSERT INTO rounds (id, user_id, game_id)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, game_id) DO UPDATE SET updated_at = NOW()
		RETURNING id, user_id, game_id, hint, status, updated_at, created_at
	`, id, userID, gameID).Scan(
		&round.ID,
		&round.UserID,
		&round.GameID,
		&round.Hint,
		&round.Status,
		&round.UpdatedAt,
		&round.CreatedAt,
	)
	if err != nil {
		log.Error().Err(err).
			Str("user_id", userID).
			Str("game_id", gameID).
			Msgf("database.postgres.dbClient.CreateRound: failed to create round: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.CreateRound: failed to create round: %v", err.Error()))
	}

	return round, nil
}

func (d *dbClient) GetRoundByUserIDForGame(ctx context.Context, userID string, gameID string) (*entities_round_v1.Round, error) {
	round := &entities_round_v1.Round{}

	err := d.connection.DB.QueryRowContext(ctx, `
		SELECT id, user_id, game_id, hint, status, has_won, updated_at, created_at
		FROM rounds
		WHERE user_id = $1 AND game_id = $2
	`, userID, gameID).Scan(
		&round.ID,
		&round.UserID,
		&round.GameID,
		&round.Hint,
		&round.Status,
		&round.HasWon,
		&round.UpdatedAt,
		&round.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).
				Str("user_id", userID).
				Msgf("database.postgres.dbClient.GetRoundByUserIDForGame: round not found")
			return nil, errors.NewNotFoundError("database.postgres.dbClient.GetRoundByUserIDForGame: round not found")
		}

		log.Error().Err(err).
			Str("user_id", userID).
			Str("game_id", gameID).
			Msgf("database.postgres.dbClient.GetRoundByUserIDForGame: failed to get round: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.GetRoundByUserIDForGame: failed to get round: %v", err.Error()))
	}

	return round, nil
}

func (d *dbClient) StartRound(ctx context.Context, userID string, gameID string) error {
	_, err := d.connection.DB.ExecContext(ctx, `
		UPDATE rounds
		SET status = $3
		WHERE user_id = $1 AND game_id = $2
	`, userID, gameID, constants.RoundStatusStarted)
	if err != nil {
		log.Error().Err(err).
			Str("user_id", userID).
			Str("game_id", gameID).
			Msgf("database.postgres.dbClient.StartRound: failed to start round: %v", err.Error())
		return errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.StartRound: failed to start round: %v", err.Error()))
	}

	return nil
}

func (d *dbClient) FinishRound(ctx context.Context, userID string, gameID string, hasWon bool) error {
	_, err := d.connection.DB.ExecContext(ctx, `
		UPDATE rounds
		SET status = $4, has_won = $3
		WHERE user_id = $1 AND game_id = $2
	`, userID, gameID, hasWon, constants.RoundStatusFinished)
	if err != nil {
		log.Error().Err(err).
			Str("user_id", userID).
			Str("game_id", gameID).
			Msgf("database.postgres.dbClient.FinishRound: failed to finish round: %v", err.Error())
		return errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.FinishRound: failed to finish round: %v", err.Error()))
	}

	return nil
}

func (d *dbClient) UpdateRound(ctx context.Context, userID string, gameID string) (*entities_round_v1.Round, error) {
	round := &entities_round_v1.Round{}

	err := d.connection.DB.QueryRowContext(ctx, `
		UPDATE rounds
		SET hint = rounds.hint + 1
		WHERE user_id = $1 AND game_id = $2
		RETURNING id, user_id, game_id, hint, status, updated_at, created_at
	`, userID, gameID).Scan(
		&round.ID,
		&round.UserID,
		&round.GameID,
		&round.Hint,
		&round.Status,
		&round.UpdatedAt,
		&round.CreatedAt,
	)
	if err != nil {
		log.Error().Err(err).
			Str("user_id", userID).
			Str("game_id", gameID).
			Msgf("database.postgres.dbClient.UpdateRound: failed to update round: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.UpdateRound: failed to update round: %v", err.Error()))
	}

	return round, nil
}
