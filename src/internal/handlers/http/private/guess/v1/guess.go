package entities_guess_v1

type Guess struct {
	GameID string `json:"game_id"`
	Guess  string `json:"guess"`
}
