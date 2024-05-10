package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/rainycape/unidecode"
	"github.com/rs/zerolog/log"
	database_postgres "github.com/teyz/songify-svc/internal/database/postgres"
	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
)

type Songs struct {
	Songs []*entities_song_v1.Song_Create `json:"songs"`
}

func main() {
	file, err := os.Open("songs.json")
	if err != nil {
		log.Error().Err(err).Msg("seeds.Run: error on open file")
	}
	defer file.Close()

	var songs *Songs
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&songs); err != nil {
		log.Error().Err(err).Msg("seeds.Run: error on decode file")
	}

	dbHost := "127.0.0.1"
	dbPort := "5432"
	dbName := "songify-db"
	dbUser := "root"
	dbPassword := "root"

	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPassword)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error().Err(err).Msg("seeds.Run: error on open db connection")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Error().Err(err).Msg("seeds.Run: error on ping db")
	}

	for _, song := range songs.Songs {
		id := database_postgres.GenerateDataPrefixWithULID(database_postgres.Song)
		title := strings.ReplaceAll(unidecode.Unidecode(strings.ToLower(song.Title)), " ", "")
		artist := strings.ReplaceAll(unidecode.Unidecode(strings.ToLower(song.Artist)), " ", "")
		db.ExecContext(context.Background(), `
			INSERT INTO songs (id, title, artist, artist_image_url, lyrics, image_url, released_year, musical_style)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`, id, title, artist, song.ArtistImageURL, song.Lyrics, song.ImageURL, song.ReleasedYear, song.MusicalStyle)
	}
}
