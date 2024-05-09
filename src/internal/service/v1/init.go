package service_v1

import (
	"context"

	"github.com/teyz/songify-svc/internal/database"
)

type Service struct {
	store database.Database
}

func NewService(ctx context.Context, store database.Database) (*Service, error) {
	return &Service{
		store: store,
	}, nil
}
