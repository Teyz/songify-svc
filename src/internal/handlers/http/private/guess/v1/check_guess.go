package handlers_http_private_guess_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	entities_guess_v1 "github.com/teyz/songify-svc/internal/entities/guess/v1"
	pkg_http "github.com/teyz/songify-svc/internal/pkg/http"
)

type CheckGuessResponse struct {
	IsCorrect bool `json:"is_correct"`
}

func (h *Handler) CheckGuess(c echo.Context) error {
	ctx := c.Request().Context()

	songID := c.Param("song_id")
	if songID == "" {
		log.Error().Msg("handlers.http.private.guess.v1.check_guess.Handler.CheckGuess: can not get song_id from context")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	artist := c.QueryParam("artist")
	if artist == "" {
		log.Error().Msg("handlers.http.private.guess.v1.check_guess.Handler.CheckGuess: can not get artist from context")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	title := c.QueryParam("title")
	if title == "" {
		log.Error().Msg("handlers.http.private.guess.v1.check_guess.Handler.CheckGuess: can not get title from context")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	isCorrect, err := h.service.CheckGuess(ctx, &entities_guess_v1.Guess{
		SongID: songID,
		Artist: artist,
		Title:  title,
	})
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, CheckGuessResponse{
		IsCorrect: isCorrect,
	}))
}
