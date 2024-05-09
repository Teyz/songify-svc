package database

import (
	"context"

	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
)

type Database interface {
	GetSongByID(ctx context.Context, id string) (*entities_song_v1.Song, error)
	GetRandomSong(ctx context.Context) (*entities_song_v1.Song, error)
}
