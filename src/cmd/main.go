package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/teyz/songify-svc/internal/config"
	"github.com/teyz/songify-svc/internal/cron"
	database_postgres "github.com/teyz/songify-svc/internal/database/postgres"
	handlers_http "github.com/teyz/songify-svc/internal/handlers/http"
	pkg_config "github.com/teyz/songify-svc/internal/pkg/config"
	pkg_postgres "github.com/teyz/songify-svc/internal/pkg/database/postgres"
	service "github.com/teyz/songify-svc/internal/service/v1"
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

	service, err := service.NewService(ctx, databaseClient)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create user store service")
	}

	newCron, err := cron.NewCron(ctx, service)
	if err != nil {
		log.Fatal().Err(err).
			Msg("main: unable to create cron service")
	}

	newCron.CreateGameCron(ctx)

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
