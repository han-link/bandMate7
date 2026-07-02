package service

import (
	"bandMate7/internal/model"
	"bandMate7/internal/store"
	"context"
)

type Performances interface {
	CreatePerformance(ctx context.Context, input CreatePerformanceRequest) (*model.Performance, error)
}

type Services struct {
	Performances Performances
}

func NewServices(store *store.Storage) Services {
	return Services{
		Performances: &PerformanceService{store},
	}
}
