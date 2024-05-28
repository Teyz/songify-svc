package service_v1

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	entities_round_v1 "github.com/teyz/songify-svc/internal/entities/round/v1"
)

func (s *Service) GetRoundByUserIDForGame(ctx context.Context, userID string, gameID string) (*entities_round_v1.Round, error) {
	key := generateRoundCacheKeyWithUserIDAndGameID(userID, gameID)

	cachedRound, err := s.cache.Get(ctx, key)
	if err == nil {
		var round *entities_round_v1.Round
		err = json.Unmarshal([]byte(cachedRound), &round)
		if err != nil {
			log.Error().Err(err).
				Msg("service.v1.service.GetRoundByUserIDForGame: unable to unmarshal round")
		} else {
			return round, nil
		}
	}

	round, err := s.store.GetRoundByUserIDForGame(ctx, userID, gameID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(round)
	if err != nil {
		log.Error().Err(err).
			Msg("service.v1.service.GetRoundByUserIDForGame: unable to marshal round")
	} else {
		s.cache.SetEx(ctx, key, bytes, roundCacheDuration)
	}

	return round, nil
}
