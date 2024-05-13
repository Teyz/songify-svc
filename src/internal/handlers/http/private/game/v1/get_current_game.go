package handlers_http_private_game_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	entities_game_v1 "github.com/teyz/songify-svc/internal/entities/game/v1"
	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
	pkg_http "github.com/teyz/songify-svc/internal/pkg/http"
)

type GetCurrentGameResponse struct {
	Game *entities_game_v1.Game `json:"game"`
	Song *entities_song_v1.Song `json:"song"`
}

func (h *Handler) GetCurrentGame(c echo.Context) error {
	ctx := c.Request().Context()

	currentGame, err := h.service.GetCurrentGame(ctx)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	song, err := h.service.GetSongByID(ctx, currentGame.SongID)
	if err != nil {
		return c.JSON(pkg_http.TranslateError(ctx, err))
	}

	return c.JSON(http.StatusOK, pkg_http.NewHTTPResponse(http.StatusOK, pkg_http.MessageSuccess, GetCurrentGameResponse{
		Game: &entities_game_v1.Game{
			ID:        currentGame.ID,
			SongID:    currentGame.SongID,
			CreatedAt: currentGame.CreatedAt,
		},
		Song: &entities_song_v1.Song{
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
