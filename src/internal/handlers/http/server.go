package handlers_http

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/teyz/songify-svc/internal/handlers"
	handlers_http_private_game_v1 "github.com/teyz/songify-svc/internal/handlers/http/private/game/v1"
	handlers_http_private_guess_v1 "github.com/teyz/songify-svc/internal/handlers/http/private/guess/v1"
	handlers_http_private_health_v1 "github.com/teyz/songify-svc/internal/handlers/http/private/health/v1"
	handlers_http_private_song_v1 "github.com/teyz/songify-svc/internal/handlers/http/private/song/v1"
	handlers_http_private_user_v1 "github.com/teyz/songify-svc/internal/handlers/http/private/user/v1"
	pkg_http "github.com/teyz/songify-svc/internal/pkg/http"
	service_v1 "github.com/teyz/songify-svc/internal/service/v1"
)

type httpServer struct {
	router  *echo.Echo
	config  pkg_http.HTTPServerConfig
	service *service_v1.Service
}

func NewServer(ctx context.Context, cfg pkg_http.HTTPServerConfig, service *service_v1.Service) (handlers.Server, error) {
	return &httpServer{
		router:  echo.New(),
		config:  cfg,
		service: service,
	}, nil
}

func (s *httpServer) Setup(ctx context.Context) error {
	log.Info().
		Msg("handlers.http.httpServer.Setup: Setting up HTTP server...")

	// setup handlers
	privateSongsV1Handlers := handlers_http_private_song_v1.NewHandler(ctx, s.service)
	privateGamesV1Handlers := handlers_http_private_game_v1.NewHandler(ctx, s.service)
	privateGuessV1Handlers := handlers_http_private_guess_v1.NewHandler(ctx, s.service)
	privateHealhV1Handlers := handlers_http_private_health_v1.NewHandler(ctx, s.service)
	privateUserV1Handlers := handlers_http_private_user_v1.NewHandler(ctx, s.service)

	// setup middlewares
	//s.router.Use(middleware.Logger())
	s.router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.CORS())

	// health endpoints
	s.router.GET("/health", privateHealhV1Handlers.HealthCheck)

	// private endpoints
	privateV1 := s.router.Group("/private/v1")

	// songs endpoints
	songsV1 := privateV1.Group("/songs")
	songsV1.GET("/:id", privateSongsV1Handlers.GetSongByID)
	songsV1.GET("/", privateSongsV1Handlers.GetRandomSong)

	// games endpoints
	gamesV1 := privateV1.Group("/games")
	gamesV1.POST("/", privateGamesV1Handlers.CreateGame)
	gamesV1.GET("/", privateGamesV1Handlers.GetCurrentGame)

	// guess endpoints
	guessV1 := privateV1.Group("/guess")
	guessV1.GET("/:song_id", privateGuessV1Handlers.CheckGuess)

	// user endpoints
	userV1 := privateV1.Group("/users")
	userV1.POST("/", privateUserV1Handlers.CreateUser)
	userV1.GET("/:id", privateUserV1Handlers.GetUserByID)

	return nil
}

func (s *httpServer) Start(ctx context.Context) error {
	log.Info().
		Uint16("port", s.config.Port).
		Msg("handlers.http.httpServer.Start: Starting HTTP server...")

	return s.router.Start(fmt.Sprintf(":%d", s.config.Port))
}

func (s *httpServer) Stop(ctx context.Context) error {
	log.Info().
		Msg("handlers.http.httpServer.Stop: Stopping HTTP server...")

	return s.router.Shutdown(ctx)
}
