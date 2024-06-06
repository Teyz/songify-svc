package database_postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rs/zerolog/log"

	entities_user_v1 "github.com/teyz/songify-svc/internal/entities/user/v1"
	"github.com/teyz/songify-svc/pkg/errors"
)

func (d *dbClient) CreateUser(ctx context.Context) (*entities_user_v1.User, error) {
	user := &entities_user_v1.User{}

	id := GenerateDataPrefixWithULID(User)

	err := d.connection.DB.QueryRowContext(ctx,
		`
			INSERT INTO users (id)
			VALUES ($1)
			RETURNING id, created_at
		`,
		id).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		log.Error().Err(err).
			Str("id", id).
			Msgf("database.postgres.dbClient.CreateUser: failed to create user: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.CreateUser: failed to create user: %v", err.Error()))
	}

	return user, nil
}

func (d *dbClient) GetUserByID(ctx context.Context, id string) (*entities_user_v1.User, error) {
	user := &entities_user_v1.User{}

	err := d.connection.DB.QueryRowContext(ctx,
		`
			SELECT id, created_at
			FROM users
			WHERE id = $1
		`,
		id).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).
				Str("user_id", id).
				Msgf("database.postgres.dbClient.GetUserByID: user not found")
			return nil, errors.NewNotFoundError("database.postgres.dbClient.GetUserByID: user not found")
		}

		log.Error().Err(err).
			Str("id", id).
			Msgf("database.postgres.dbClient.GetUserByID: failed to get user: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.GetUserByID: failed to get user: %v", err.Error()))
	}

	return user, nil
}
