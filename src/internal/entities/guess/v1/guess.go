package entities_guess_v1

type Guess struct {
	GameID          string `json:"game_id"`
	Artist          string `json:"artist"`
	IsArtistCorrect bool   `json:"is_artist_correct"`
	Title           string `json:"title"`
	IsTitleCorrect  bool   `json:"is_title_correct"`
}

type Guesses struct {
	IsArtistCorrect bool     `json:"is_artist_correct"`
	IsTitleCorrect  bool     `json:"is_title_correct"`
	Guesses         []*Guess `json:"guesses"`
}
