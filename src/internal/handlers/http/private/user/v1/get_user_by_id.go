package handlers_http_private_user_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	entities_user_v1 "github.com/teyz/songify-svc/internal/entities/user/v1"
	pkg_http "github.com/teyz/songify-svc/internal/pkg/http"
)

type GetUserByIDResponse struct {
	User *entities_user_v1.User `json:"user"`
}

func (h *Handler) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		log.Error().Msg("handlers.http.private.user.v1.get_user_by_id.Handler.GetUserByID: can not get id from context")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	user, err := h.service.GetUserByID(ctx, id)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, GetUserByIDResponse{
		User: &entities_user_v1.User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
		},
	}))
}
