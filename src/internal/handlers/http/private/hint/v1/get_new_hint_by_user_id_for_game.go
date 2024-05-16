package handlers_http_private_hint_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	pkg_http "github.com/teyz/songify-svc/internal/pkg/http"
)

type GetHintByUserIDForGameResponse struct {
	HasHint  bool   `json:"has_hint"`
	HintType uint32 `json:"hint_type"`
	Hint     string `json:"hint"`
}

func (h *Handler) GetNewHintByUserIDForGame(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Param("user_id")
	if userID == "" {
		log.Error().Msg("handlers.http.private.hint.v1.get_hint_by_user_id_for_game.Handler.GetHintByUserIDForGame: can not get user_id from context")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	gameID := c.Param("game_id")
	if gameID == "" {
		log.Error().Msg("handlers.http.private.hint.v1.get_hint_by_user_id_for_game.Handler.GetHintByUserIDForGame: can not get game_id from context")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	hint, err := h.service.GetNewHintByUserIDForGame(ctx, userID, gameID)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, GetHintByUserIDForGameResponse{
		HintType: hint.HintType,
		HasHint:  hint.HintType > 0,
		Hint:     hint.Hint,
	}))
}
