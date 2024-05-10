package service_v1

import (
	"context"
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/rainycape/unidecode"
	entities_guess_v1 "github.com/teyz/songify-svc/internal/entities/guess/v1"
)

func (s *Service) CheckGuess(ctx context.Context, guess *entities_guess_v1.Guess) (bool, error) {
	song, err := s.store.CheckGuess(ctx, guess.SongID)
	if err != nil {
		return false, err
	}

	guessedTitle := strings.ReplaceAll(unidecode.Unidecode(strings.ToLower(guess.Title)), " ", "")
	guessedArtist := strings.ReplaceAll(unidecode.Unidecode(strings.ToLower(guess.Artist)), " ", "")

	titleDistance := levenshtein.ComputeDistance(song.Title, guessedTitle)
	artistDistance := levenshtein.ComputeDistance(song.Artist, guessedArtist)

	if titleDistance > 0 && artistDistance > 0 {
		return false, nil
	}

	return true, nil
}
