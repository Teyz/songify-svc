package handlers_http_private_guess_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	entities_guess_v1 "github.com/teyz/songify-svc/internal/entities/guess/v1"
	pkg_http "github.com/teyz/songify-svc/internal/pkg/http"
)

type GetGuessesByUserIDForGameResponse struct {
	IsTitleCorrect  bool                       `json:"is_title_correct"`
	IsArtistCorrect bool                       `json:"is_artist_correct"`
	Guesses         []*entities_guess_v1.Guess `json:"guesses"`
}

func (h *Handler) GetGuessesByUserIDForGame(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Param("user_id")
	if userID == "" {
		log.Error().Msg("handlers.http.private.guess.v1.get_guesses.Handler.GetGuessesByUserIDForGame: can not get user_id from context")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	gameID := c.Param("game_id")
	if gameID == "" {
		log.Error().Msg("handlers.http.private.guess.v1.get_guesses.Handler.GetGuessesByUserIDForGame: can not get game_id from context")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	guesses, err := h.service.GetGuessesByUserIDForGame(ctx, userID, gameID)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, GetGuessesByUserIDForGameResponse{
		IsTitleCorrect:  guesses.IsTitleCorrect,
		IsArtistCorrect: guesses.IsArtistCorrect,
		Guesses:         guesses.Guesses,
	}))
}
