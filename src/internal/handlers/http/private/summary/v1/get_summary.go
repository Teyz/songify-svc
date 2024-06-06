package handlers_http_private_summary_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	entities_summary_v1 "github.com/teyz/songify-svc/internal/entities/summary/v1"
	pkg_http "github.com/teyz/songify-svc/pkg/http"
)

type GetSummaryResponse struct {
	Summary *entities_summary_v1.Summary `json:"summary"`
}

func (h *Handler) GetSummary(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Param("user_id")
	if userID == "" {
		log.Error().Msg("handlers.http.private.hint.v1.get_summary.Handler.GetSummary: can not get user_id from context")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	gameID := c.Param("game_id")
	if gameID == "" {
		log.Error().Msg("handlers.http.private.hint.v1.get_summary.Handler.GetSummary: can not get game_id from context")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	summary, err := h.service.GetSummary(ctx, userID, gameID)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, GetSummaryResponse{
		Summary: summary,
	}))
}
