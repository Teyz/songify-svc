package handlers_http_private_guess_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	entities_guess_v1 "github.com/teyz/songify-svc/internal/entities/guess/v1"
	pkg_http "github.com/teyz/songify-svc/internal/pkg/http"
)

type CheckGuessRequest struct {
	UserID string `json:"user_id"`
	GameID string `json:"game_id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

type CheckGuessResponse struct {
	IsTitleCorrect  bool                       `json:"is_title_correct"`
	IsArtistCorrect bool                       `json:"is_artist_correct"`
	Guesses         []*entities_guess_v1.Guess `json:"guesses"`
}

func (h *Handler) CheckGuess(c echo.Context) error {
	ctx := c.Request().Context()

	var req CheckGuessRequest
	if err := c.Bind(&req); err != nil {
		log.Error().Err(err).Msg("handlers.http.private.guess.v1.check_guess.CheckGuess: can not bind request")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	if req.UserID == "" {
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	guesses, err := h.service.CheckGuess(ctx, req.UserID, &entities_guess_v1.Guess{
		GameID: req.GameID,
		Artist: req.Artist,
		Title:  req.Title,
	})
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, CheckGuessResponse{
		IsTitleCorrect:  guesses.IsTitleCorrect,
		IsArtistCorrect: guesses.IsArtistCorrect,
		Guesses:         guesses.Guesses,
	}))
}
