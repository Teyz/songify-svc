package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/teyz/songify-svc/internal/pkg/cache"
)

type cacheClient struct {
	rdb redis.UniversalClient
}

func NewRedisCache(ctx context.Context, rdb redis.UniversalClient) cache.Cache {
	return &cacheClient{rdb: rdb}
}
