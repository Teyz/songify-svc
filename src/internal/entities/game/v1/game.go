package entities_game_v1

import "time"

type Game struct {
	ID        string    `json:"id"`
	SongID    string    `json:"song_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CurrentGame struct {
	ID        string    `json:"id"`
	SongID    string    `json:"song_id"`
	Lyrics    []string  `json:"lyrics"`
	CreatedAt time.Time `json:"created_at"`
}
