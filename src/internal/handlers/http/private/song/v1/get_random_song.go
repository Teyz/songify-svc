package handlers_http_private_song_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"

	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
	pkg_http "github.com/teyz/songify-svc/pkg/http"
)

type GetRandomSongResponse struct {
	Song *entities_song_v1.Song_HTTP `json:"song"`
}

func (h *Handler) GetRandomSong(c echo.Context) error {
	ctx := c.Request().Context()

	song, err := h.service.GetRandomSong(ctx)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, GetRandomSongResponse{
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
