package service_v1

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/rainycape/unidecode"
	"github.com/rs/zerolog/log"
	entities_guess_v1 "github.com/teyz/songify-svc/internal/entities/guess/v1"
	"github.com/teyz/songify-svc/internal/pkg/errors"
)

func (s *Service) CheckGuess(ctx context.Context, userID string, guess *entities_guess_v1.Guess) (*entities_guess_v1.Guesses, error) {
	doesUserCanGuess, err := s.store.CheckIfUserCanGuess(ctx, userID, guess.GameID)
	if err != nil {
		return nil, err
	}

	if !doesUserCanGuess {
		err := s.store.FinishRound(ctx, userID, guess.GameID, false)
		if err != nil {
			return nil, err
		}
		return nil, errors.NewBadRequestError("user can not guess anymore")
	}

	s.cache.Del(ctx, generateGuessCacheKeyWithUserIDAndGameID(userID, guess.GameID))

	game, err := s.store.GetCurrentGame(ctx)
	if err != nil {
		return nil, err
	}

	song, err := s.store.CheckGuess(ctx, game.SongID)
	if err != nil {
		return nil, err
	}

	guessedTitle := strings.ReplaceAll(unidecode.Unidecode(strings.ToLower(guess.Title)), " ", "")
	guessedArtist := strings.ReplaceAll(unidecode.Unidecode(strings.ToLower(guess.Artist)), " ", "")

	titleDistance := levenshtein.ComputeDistance(song.Title, guessedTitle)
	artistDistance := levenshtein.ComputeDistance(song.Artist, guessedArtist)

	err = s.store.AddGuess(ctx, userID, &entities_guess_v1.Guess{
		GameID:          guess.GameID,
		Artist:          guess.Artist,
		Title:           guess.Title,
		IsTitleCorrect:  titleDistance == 0,
		IsArtistCorrect: artistDistance == 0,
	})
	if err != nil {
		return nil, err
	}

	guesses, err := s.store.GetGuessesByUserIDForGame(ctx, userID, guess.GameID)
	if err != nil {
		return nil, err
	}

	if len(guesses.Guesses) == 3 {
		err := s.store.FinishRound(ctx, userID, guess.GameID, (titleDistance == 0 && artistDistance == 0))
		if err != nil {
			return nil, err
		}
	}

	if titleDistance > 0 && artistDistance > 0 {
		return &entities_guess_v1.Guesses{
			IsTitleCorrect:  false,
			IsArtistCorrect: false,
			Guesses:         guesses.Guesses,
		}, nil
	}

	if titleDistance == 0 && artistDistance == 0 {
		err := s.store.FinishRound(ctx, userID, guess.GameID, true)
		if err != nil {
			return nil, err
		}
	}

	return &entities_guess_v1.Guesses{
		IsTitleCorrect:  titleDistance == 0,
		IsArtistCorrect: artistDistance == 0,
		Guesses:         guesses.Guesses,
	}, nil
}

func (s *Service) GetGuessesByUserIDForGame(ctx context.Context, userID string, gameID string) (*entities_guess_v1.Guesses, error) {
	key := generateGuessCacheKeyWithUserIDAndGameID(userID, gameID)

	cachedGuesses, err := s.cache.Get(ctx, key)
	if err == nil {
		var guesses *entities_guess_v1.Guesses
		err = json.Unmarshal([]byte(cachedGuesses), &guesses)
		if err != nil {
			log.Error().Err(err).
				Msg("service.v1.service.GetGuessesByUserIDForGame: unable to unmarshal guesses")
		} else {
			return guesses, nil
		}
	}

	guesses, err := s.store.GetGuessesByUserIDForGame(ctx, userID, gameID)
	if err != nil {
		return nil, err
	}

	if len(guesses.Guesses) == 0 {
		_, err := s.store.CreateRound(ctx, userID, gameID)
		if err != nil {
			return nil, err
		}
	}

	bytes, err := json.Marshal(guesses)
	if err != nil {
		log.Error().Err(err).
			Msg("service.v1.service.GetGuessesByUserIDForGame: unable to marshal guesses")
	} else {
		s.cache.SetEx(ctx, key, bytes, gameCacheDuration)
	}

	return guesses, nil
}
