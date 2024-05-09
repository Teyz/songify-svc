package entities_song_v1

import "time"

type Song struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Artist         string    `json:"artist"`
	ArtistImageURL string    `json:"artist_image_url"`
	Lyrics         string    `json:"lyrics"`
	ImageURL       string    `json:"image_url"`
	ReleasedYear   int       `json:"released_year"`
	MusicalStyle   string    `json:"musical_style"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Song_HTTP struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	Artist         string    `json:"artist"`
	ArtistImageURL string    `json:"artist_image_url"`
	Lyrics         string    `json:"lyrics"`
	ImageURL       string    `json:"image_url"`
	ReleasedYear   int       `json:"released_year"`
	MusicalStyle   string    `json:"musical_style"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Song_Create struct {
	Title          string `json:"title"`
	Artist         string `json:"artist"`
	ArtistImageURL string `json:"artist_image_url"`
	Lyrics         string `json:"lyrics"`
	ImageURL       string `json:"image_url"`
	ReleasedYear   int    `json:"released_year"`
	MusicalStyle   string `json:"musical_style"`
}
