package handlers_http

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"

	"github.com/teyz/songify-svc/internal/handlers"
	handlers_http_private_song_v1 "github.com/teyz/songify-svc/internal/handlers/http/private/song/v1"
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

	// setup middlewares
	s.router.Use(middleware.Logger())
	s.router.Use(middleware.Recover())
	s.router.Use(middleware.CORS())

	// private endpoints
	privateV1 := s.router.Group("/private/v1")

	// songs endpoints
	songsV1 := privateV1.Group("/songs")
	songsV1.GET("/:id", privateSongsV1Handlers.GetSongByID)
	songsV1.GET("/", privateSongsV1Handlers.GetRandomSong)

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
