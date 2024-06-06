package config

import (
	pkg_redis "github.com/teyz/songify-svc/pkg/cache/redis"
	pkg_postgres "github.com/teyz/songify-svc/pkg/database/postgres"
	pkg_http "github.com/teyz/songify-svc/pkg/http"
)

type Config struct {
	ServiceName string `env:"SERVICE_NAME" envDefault:"songify-svc"`
	Environment string `env:"ENVIRONMENT" envDefault:"local"`

	HTTPServerConfig pkg_http.HTTPServerConfig
	PostgresConfig   pkg_postgres.PostgresConfig
	RedisConfig      pkg_redis.RedisConfig
}
