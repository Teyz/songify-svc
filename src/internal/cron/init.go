package cron

import (
	"context"

	service_v1 "github.com/teyz/songify-svc/internal/service/v1"
)

type cron struct {
	service *service_v1.Service
}

func NewCron(ctx context.Context, service *service_v1.Service) (*cron, error) {
	return &cron{
		service: service,
	}, nil
}
