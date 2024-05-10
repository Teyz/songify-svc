package database

import (
	"context"

	entities_game_v1 "github.com/teyz/songify-svc/internal/entities/game/v1"
	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
)

type Database interface {
	GetSongByID(ctx context.Context, id string) (*entities_song_v1.Song, error)
	GetRandomSong(ctx context.Context) (*entities_song_v1.Song, error)

	CreateGame(ctx context.Context, songID string) (*entities_game_v1.Game, error)
	GetCurrentGame(ctx context.Context) (*entities_game_v1.Game, error)

	CheckGuess(ctx context.Context, songID string) (*entities_song_v1.Song_Guess, error)
}
