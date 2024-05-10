package entities_guess_v1

type Guess struct {
	SongID string `json:"song_id"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}
