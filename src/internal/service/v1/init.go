package service_v1

import (
	"context"
	"fmt"
	"time"

	"github.com/teyz/songify-svc/internal/database"
	pkg_cache "github.com/teyz/songify-svc/pkg/cache"
)

const (
	gameCacheDuration    = time.Hour * 2
	guessCacheDuration   = time.Hour * 2
	songCacheDuration    = time.Hour * 24 * 30
	summaryCacheDuration = time.Hour * 2
	roundCacheDuration   = time.Hour * 2
	userCacheDuration    = time.Hour * 24
)

func generateGameCacheKey() string {
	return "songify-svc:game"
}

func generateUserCacheKeyWithID(userID string) string {
	return fmt.Sprintf("songify-svc:user:user_id:%v", userID)
}

func generateSongCacheKeyWithID(songID string) string {
	return fmt.Sprintf("songify-svc:song:song_id:%v", songID)
}

func generateGuessCacheKeyWithUserIDAndGameID(userID string, gameID string) string {
	return fmt.Sprintf("songify-svc:guess:user_id:%v:game_id:%v", userID, gameID)
}

func generateSummaryCacheKeyWithUserIDAndGameID(userID string, gameID string) string {
	return fmt.Sprintf("songify-svc:summary:user_id:%v:game_id:%v", userID, gameID)
}

func generateRoundCacheKeyWithUserIDAndGameID(userID string, gameID string) string {
	return fmt.Sprintf("songify-svc:round:user_id:%v:game_id:%v", userID, gameID)
}

type Service struct {
	store database.Database
	cache pkg_cache.Cache
}

func NewService(ctx context.Context, store database.Database, cache pkg_cache.Cache) (*Service, error) {
	return &Service{
		store: store,
		cache: cache,
	}, nil
}
