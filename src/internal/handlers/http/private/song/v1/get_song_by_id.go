package handlers_http_private_song_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"

	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
	pkg_http "github.com/teyz/songify-svc/pkg/http"
)

type GetSongByIDResponse struct {
	Song *entities_song_v1.Song_HTTP `json:"song"`
}

func (h *Handler) GetSongByID(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	if id == "" {
		log.Error().Msg("handlers.http.private.song.v1.get_song_by_id.Handler.GetSongByID: can not get id from context")
		return c.JSON(http.StatusBadRequest, pkg_http.NewHTTPResponse(http.StatusBadRequest, pkg_http.MessageBadRequestError, nil))
	}

	song, err := h.service.GetSongByID(ctx, id)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, GetSongByIDResponse{
		Song: &entities_song_v1.Song_HTTP{
			ID:             song.ID,
			Title:          song.Title,
			Artist:         song.Artist,
			ArtistImageURL: song.ArtistImageURL,
			Lyrics:         song.Lyrics,
			ImageURL:       song.ImageURL,
			ReleasedYear:   song.ReleasedYear,
			MusicalStyle:   song.MusicalStyle,
			CreatedAt:      song.CreatedAt,
			UpdatedAt:      song.UpdatedAt,
		},
	}))
}
