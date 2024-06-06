package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/teyz/songify-svc/internal/config"
	database_postgres "github.com/teyz/songify-svc/internal/database/postgres"
	handlers_http "github.com/teyz/songify-svc/internal/handlers/http"
	service "github.com/teyz/songify-svc/internal/service/v1"
	pkg_redis "github.com/teyz/songify-svc/pkg/cache/redis"
	pkg_config "github.com/teyz/songify-svc/pkg/config"
	pkg_postgres "github.com/teyz/songify-svc/pkg/database/postgres"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	cfg := &config.Config{}
	err := pkg_config.ParseConfig(cfg)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to parse config")
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	databaseConnection, err := pkg_postgres.NewDatabaseConnection(ctx, &cfg.PostgresConfig)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create database connection")
	}
	databaseClient := database_postgres.NewClient(ctx, databaseConnection)

	cacheConnection := pkg_redis.GetConnection(ctx, &cfg.RedisConfig)
	redisClient := pkg_redis.NewRedisCache(ctx, cacheConnection)

	service, err := service.NewService(ctx, databaseClient, redisClient)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create user store service")
	}

	// create http server
	httpServer, err := handlers_http.NewServer(ctx, cfg.HTTPServerConfig, service)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create http server")
	}

	// setup http server
	if err := httpServer.Setup(ctx); err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to setup http server")
	}

	// start http server
	if err := httpServer.Start(ctx); err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to start http server")
	}

	<-sigs
	cancel()

	// stop http server
	if err := httpServer.Stop(ctx); err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to stop http server")
	}

	os.Exit(0)
}
