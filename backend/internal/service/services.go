package service

import (
	"bandMate7/internal/model"
	"bandMate7/internal/store"
	"context"
)

type Performances interface {
	Create(ctx context.Context, input CreatePerformanceRequest) (*model.Performance, error)
	GetAll(ctx context.Context) (*[]model.Performance, error)
}

type Services struct {
	Performances Performances
}

func NewServices(store *store.Storage, baseUrl string) Services {
	return Services{
		Performances: &PerformanceService{baseUrl, store},
	}
}
