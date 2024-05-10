package service_v1

import (
	"context"

	entities_game_v1 "github.com/teyz/songify-svc/internal/entities/game/v1"
)

func (s *Service) CreateGame(ctx context.Context) (*entities_game_v1.Game, error) {
	song, err := s.store.GetRandomSong(ctx)
	if err != nil {
		return nil, err
	}

	game, err := s.store.CreateGame(ctx, song.ID)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (s *Service) GetCurrentGame(ctx context.Context) (*entities_game_v1.Game, error) {
	game, err := s.store.GetCurrentGame(ctx)
	if err != nil {
		return nil, err
	}

	return game, nil
}
