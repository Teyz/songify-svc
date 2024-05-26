package entities_summary_v1

import (
	entities_round_v1 "github.com/teyz/songify-svc/internal/entities/round/v1"
	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
)

type Summary struct {
	Song  *entities_song_v1.Song_Lite   `json:"song"`
	Round *entities_round_v1.Round_Lite `json:"round"`
}
