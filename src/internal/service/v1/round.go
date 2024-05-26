package service_v1

import (
	"context"

	entities_round_v1 "github.com/teyz/songify-svc/internal/entities/round/v1"
)

func (s *Service) GetRoundByUserIDForGame(ctx context.Context, userID string, gameID string) (*entities_round_v1.Round, error) {
	round, err := s.store.GetRoundByUserIDForGame(ctx, userID, gameID)
	if err != nil {
		return nil, err
	}

	return round, nil
}
