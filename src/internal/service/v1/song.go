package service_v1

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
)

func (s *Service) GetSongByID(ctx context.Context, id string) (*entities_song_v1.Song, error) {
	key := generateSongCacheKeyWithID(id)

	cachedSong, err := s.cache.Get(ctx, key)
	if err == nil {
		var song *entities_song_v1.Song
		err = json.Unmarshal([]byte(cachedSong), &song)
		if err != nil {
			log.Error().Err(err).
				Str("song_id", id).
				Msg("service.v1.service.GetSongByID: unable to unmarshal song")
		} else {
			return song, nil
		}
	}

	song, err := s.store.GetSongByID(ctx, id)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(song)
	if err != nil {
		log.Error().Err(err).
			Str("song_id", id).
			Msg("service.v1.service.GetSongByID: unable to marshal song")
	} else {
		s.cache.SetEx(ctx, key, bytes, songCacheDuration)
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
