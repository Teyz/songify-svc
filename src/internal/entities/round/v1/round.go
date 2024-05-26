package entities_round_v1

import "time"

type Round struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	GameID    string    `json:"game_id"`
	Hint      uint16    `json:"hint"`
	Status    string    `json:"status"`
	HasWon    bool      `json:"has_won"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Round_Lite struct {
	Hint   uint16 `json:"hint"`
	Status string `json:"status"`
	HasWon bool   `json:"has_won"`
}
