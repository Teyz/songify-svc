package entities_game_v1

import "time"

type Game struct {
	ID        string    `json:"id"`
	SongID    string    `json:"song_id"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}