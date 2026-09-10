package service

import (
	"bandMate7/internal/model"
	"bandMate7/internal/store"
	"context"
	"net/http"
)

type Performances interface {
	Create(ctx context.Context, input CreatePerformanceRequest) (*model.Performance, error)
	GetAll(ctx context.Context, r *http.Request) (*[]model.Performance, error)
}

type Setlists interface {
	Create(ctx context.Context, payload CreateSetlistRequest) (*model.Setlist, error)
	GetAll(ctx context.Context) ([]SelistWithoutPerformance, error)
	ChangeOrder(ctx context.Context, payload ChangeOrderPayload, setlist *model.Setlist) (*model.Setlist, error)
}

type Services struct {
	Performances Performances
	Setlists     Setlists
}

func NewServices(store *store.Storage, baseUrl string) Services {
	return Services{
		Performances: &PerformanceService{baseUrl, store},
		Setlists:     &SetlistsService{store},
	}
}
