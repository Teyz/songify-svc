package service_v1

import (
	"context"
	"fmt"

	entities_hint_v1 "github.com/teyz/songify-svc/internal/entities/hint/v1"
	"github.com/teyz/songify-svc/pkg/errors"
)

func (s *Service) GetNewHintByUserIDForGame(ctx context.Context, userID string, gameID string) (*entities_hint_v1.Hint, error) {
	round, err := s.store.GetRoundByUserIDForGame(ctx, userID, gameID)
	if err != nil {
		return nil, err
	}

	if round.Hint >= 3 {
		return nil, errors.NewBadRequestError("user can not ask for hint anymore")
	}

	_, err = s.store.UpdateRound(ctx, userID, gameID)
	if err != nil {
		return nil, err
	}

	round, err = s.store.GetRoundByUserIDForGame(ctx, userID, gameID)
	if err != nil {
		return nil, err
	}

	game, err := s.store.GetCurrentGame(ctx)
	if err != nil {
		return nil, err
	}

	song, err := s.store.GetSongByID(ctx, game.SongID)
	if err != nil {
		return nil, err
	}

	hint := &entities_hint_v1.Hint{
		HintType: uint32(round.Hint),
	}

	if round.Hint == 1 {
		hint.Hint = song.MusicalStyle
	} else if round.Hint == 2 {
		hint.Hint = fmt.Sprintf("%d", song.ReleasedYear)
	} else {
		hint.Hint = song.ArtistImageURL
	}

	return hint, nil
}
