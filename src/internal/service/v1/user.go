package service_v1

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"

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
	key := generateUserCacheKeyWithID(userID)

	cachedUser, err := s.cache.Get(ctx, key)
	if err == nil {
		var user *entities_user_v1.User
		err = json.Unmarshal([]byte(cachedUser), &user)
		if err != nil {
			log.Error().Err(err).
				Msg("service.v1.service.GetUserByID: unable to unmarshal user")
		} else {
			return user, nil
		}
	}

	user, err := s.store.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(user)
	if err != nil {
		log.Error().Err(err).
			Msg("service.v1.service.GetUserByID: unable to marshal user")
	} else {
		s.cache.SetEx(ctx, key, bytes, userCacheDuration)
	}

	return user, nil
}
