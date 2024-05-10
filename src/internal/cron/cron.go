package cron

import (
	"context"

	robfig_cron "github.com/robfig/cron"
	"github.com/rs/zerolog/log"
)

func (cron *cron) CreateGameCron(ctx context.Context) {
	c := robfig_cron.New()

	c.AddFunc("@midnight", func() {
		_, err := cron.service.CreateGame(ctx)
		if err != nil {
			log.Info().Msg("cron.cron.CreateGameCron: game created")
		}
	})

	c.Start()
}
