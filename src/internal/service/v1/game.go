package service_v1

import (
	"context"
	"strings"

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

func (s *Service) GetCurrentGame(ctx context.Context) (*entities_game_v1.CurrentGame, error) {
	game, err := s.store.GetCurrentGame(ctx)
	if err != nil {
		return nil, err
	}

	song, err := s.store.GetSongByID(ctx, game.SongID)
	if err != nil {
		return nil, err
	}

	lyric := strings.Split(song.Lyrics, "\n")[0]

	return &entities_game_v1.CurrentGame{
		ID:        game.ID,
		SongID:    game.SongID,
		Lyric:     lyric,
		CreatedAt: game.CreatedAt,
	}, nil
}
