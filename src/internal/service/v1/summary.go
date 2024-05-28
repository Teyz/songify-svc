package service_v1

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	entities_round_v1 "github.com/teyz/songify-svc/internal/entities/round/v1"
	entities_song_v1 "github.com/teyz/songify-svc/internal/entities/song/v1"
	entities_summary_v1 "github.com/teyz/songify-svc/internal/entities/summary/v1"
)

func (s *Service) GetSummary(ctx context.Context, userID string, gameID string) (*entities_summary_v1.Summary, error) {
	key := generateSummaryCacheKeyWithUserIDAndGameID(userID, gameID)

	cachedSummary, err := s.cache.Get(ctx, key)
	if err == nil {
		var summary *entities_summary_v1.Summary
		err = json.Unmarshal([]byte(cachedSummary), &summary)
		if err != nil {
			log.Error().Err(err).
				Msg("service.v1.service.GetSummary: unable to unmarshal summary")
		} else {
			return summary, nil
		}
	}

	round, err := s.store.GetRoundByUserIDForGame(ctx, userID, gameID)
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

	summary := &entities_summary_v1.Summary{
		Song: &entities_song_v1.Song_Lite{
			Title:          song.Title,
			Artist:         song.Artist,
			ArtistImageURL: song.ArtistImageURL,
			Lyrics:         song.Lyrics,
			ImageURL:       song.ImageURL,
			ReleasedYear:   song.ReleasedYear,
			MusicalStyle:   song.MusicalStyle,
		},
		Round: &entities_round_v1.Round_Lite{
			Hint:   round.Hint,
			Status: round.Status,
			HasWon: round.HasWon,
		},
	}

	bytes, err := json.Marshal(summary)
	if err != nil {
		log.Error().Err(err).
			Msg("service.v1.service.GetSummary: unable to marshal summary")
	} else {
		s.cache.SetEx(ctx, key, bytes, summaryCacheDuration)
	}

	return summary, nil
}
