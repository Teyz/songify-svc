package service_v1

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/rs/zerolog/log"

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

	s.cache.Del(ctx, generateGameCacheKey())

	return game, nil
}

func (s *Service) GetCurrentGame(ctx context.Context) (*entities_game_v1.CurrentGame, error) {
	key := generateGameCacheKey()

	cacheGame, err := s.cache.Get(ctx, key)
	if err == nil {
		var game *entities_game_v1.CurrentGame
		err = json.Unmarshal([]byte(cacheGame), &game)
		if err != nil {
			log.Error().Err(err).
				Msg("service.v1.service.GetCurrentGame: unable to unmarshal game")
		} else {
			return game, nil
		}
	}

	game, err := s.store.GetCurrentGame(ctx)
	if err != nil {
		return nil, err
	}

	song, err := s.store.GetSongByID(ctx, game.SongID)
	if err != nil {
		return nil, err
	}

	lyrics := strings.Split(song.Lyrics, "\n")

	currentGame := &entities_game_v1.CurrentGame{
		ID:        game.ID,
		SongID:    game.SongID,
		Lyrics:    lyrics,
		CreatedAt: game.CreatedAt,
	}

	bytes, err := json.Marshal(currentGame)
	if err != nil {
		log.Error().Err(err).
			Msg("service.v1.service.GetCurrentGame: unable to marshal game")
	} else {
		s.cache.SetEx(ctx, key, bytes, gameCacheDuration)
	}

	return currentGame, nil
}
