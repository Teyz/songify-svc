package handlers_http_private_user_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	entities_user_v1 "github.com/teyz/songify-svc/internal/entities/user/v1"
	pkg_http "github.com/teyz/songify-svc/pkg/http"
)

type CreateUserResponse struct {
	User *entities_user_v1.User `json:"user"`
}

func (h *Handler) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()

	user, err := h.service.CreateUser(ctx)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusCreated, pkg_http.NewHTTPResponse(http.StatusCreated, pkg_http.MessageSuccess, CreateUserResponse{
		User: &entities_user_v1.User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
		},
	}))
}
