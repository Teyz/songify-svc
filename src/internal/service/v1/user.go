package service_v1

import (
	"context"

	entities_user_v1 "github.com/teyz/songify-svc/internal/entities/user/v1"
)

func (s *Service) CreateUser(ctx context.Context) (*entities_user_v1.User, error) {
	user, err := s.store.CreateUser(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUserByID(ctx context.Context, userID string) (*entities_user_v1.User, error) {
	user, err := s.store.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &entities_user_v1.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
	}, nil
}
