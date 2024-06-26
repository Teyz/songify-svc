package service_v1

import (
	"context"

	entities_game_v1 "github.com/teyz/songify-svc/internal/entities/game/v1"
	entities_guess_v1 "github.com/teyz/songify-svc/internal/entities/guess/v1"
	entities_hint_v1 "github.com/teyz/songify-svc/internal/entities/hint/v1"
	entities_round_v1 "github.com/teyz/songify-svc/internal/entities/round/v1"
	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
	entities_user_v1 "github.com/teyz/songify-svc/internal/entities/user/v1"
)

type UserStoreService interface {
	GetSongByID(ctx context.Context, id string) (*entities_song_v1.Song, error)
	GetRandomSong(ctx context.Context) (*entities_song_v1.Song, error)

	CreateGame(ctx context.Context) (*entities_game_v1.Game, error)
	GetCurrentGame(ctx context.Context) (*entities_game_v1.Game, error)

	CheckIfUserCanGuess(ctx context.Context, userID string, gameID string) (bool, error)
	CheckGuess(ctx context.Context, guess *entities_guess_v1.Guess) (bool, error)
	AddGuess(ctx context.Context, guess *entities_guess_v1.Guess) error
	GetGuessesByUserIDForGame(ctx context.Context, userID string, gameID string) (*entities_guess_v1.Guesses, error)

	CreateUser(ctx context.Context) (*entities_user_v1.User, error)
	GetUserByID(ctx context.Context, id string) (*entities_user_v1.User, error)

	CreateRound(ctx context.Context, userID string, gameID string) (*entities_round_v1.Round, error)
	GetRoundByUserIDForGame(ctx context.Context, userID string, gameID string) (*entities_round_v1.Round, error)
	GetRoundsByUserID(ctx context.Context, userID string) ([]*entities_round_v1.Round, error)
	FinishRound(ctx context.Context, userID string, gameID string, hasWon bool) error

	GetCurrentHint(ctx context.Context, userID string, gameID string) (*entities_hint_v1.Hint, error)
	GetNewHintByUserIDForGame(ctx context.Context, userID string, gameID string) (*entities_hint_v1.Hint, error)
}
