package entities_summary_v1

import (
	entities_game_v1 "github.com/teyz/songify-svc/internal/entities/game/v1"
	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
)

type Summary struct {
	Song *entities_song_v1.Song `json:"song"`
	Game *entities_game_v1.Game `json:"game"`
}
