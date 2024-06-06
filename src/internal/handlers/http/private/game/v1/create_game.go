package handlers_http_private_game_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	entities_game_v1 "github.com/teyz/songify-svc/internal/entities/game/v1"
	pkg_http "github.com/teyz/songify-svc/pkg/http"
)

type CreateGameResponse struct {
	Game *entities_game_v1.Game `json:"game"`
}

func (h *Handler) CreateGame(c echo.Context) error {
	ctx := c.Request().Context()

	game, err := h.service.CreateGame(ctx)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusCreated, pkg_http.NewHTTPResponse(http.StatusCreated, pkg_http.MessageSuccess, CreateGameResponse{
		Game: &entities_game_v1.Game{
			ID:        game.ID,
			SongID:    game.SongID,
			CreatedAt: game.CreatedAt,
		},
	}))
}
