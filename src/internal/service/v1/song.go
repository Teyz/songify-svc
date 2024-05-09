package service_v1

import (
	"context"

	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
)

func (s *Service) GetSongByID(ctx context.Context, id string) (*entities_song_v1.Song, error) {
	song, err := s.store.GetSongByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (s *Service) GetRandomSong(ctx context.Context) (*entities_song_v1.Song, error) {
	song, err := s.store.GetRandomSong(ctx)
	if err != nil {
		return nil, err
	}

	return song, nil
}
