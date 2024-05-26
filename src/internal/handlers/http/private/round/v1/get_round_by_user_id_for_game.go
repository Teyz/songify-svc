package handlers_http_private_round_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	entities_round_v1 "github.com/teyz/songify-svc/internal/entities/round/v1"
	pkg_http "github.com/teyz/songify-svc/internal/pkg/http"
)

type GetRoundByUserIDForGameResponse struct {
	Round *entities_round_v1.Round `json:"round"`
}

func (h *Handler) GetRoundByUserIDForGame(c echo.Context) error {
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

	round, err := h.service.GetRoundByUserIDForGame(ctx, userID, gameID)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, GetRoundByUserIDForGameResponse{
		Round: &entities_round_v1.Round{
			ID:        round.ID,
			GameID:    round.GameID,
			UserID:    round.UserID,
			Hint:      round.Hint,
			Status:    round.Status,
			HasWon:    round.HasWon,
			UpdatedAt: round.UpdatedAt,
			CreatedAt: round.CreatedAt,
		},
	}))
}
