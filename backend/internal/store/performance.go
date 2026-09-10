package store

import (
	"bandMate7/internal/model"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PerformanceStore struct {
	db *gorm.DB
}

func (s *PerformanceStore) GetAll(ctx context.Context, pq PaginatedQuery) ([]model.Performance, error) {
	var performances []model.Performance
	err := s.db.WithContext(ctx).
		Scopes(pq.paginateResult).
		Preload("Cover").
		Preload("Resources").
		Preload("Resources.UserRole").
		Find(&performances).
		Error
	if err != nil {
		return nil, err
	}
	return performances, nil
}

func (s *PerformanceStore) Create(ctx context.Context, performance *model.Performance) error {
	err := s.db.WithContext(ctx).
		Create(performance).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (s *PerformanceStore) GetByID(ctx context.Context, id uuid.UUID) (*model.Performance, error) {
	var performance model.Performance
	err := s.db.WithContext(ctx).Preload("Cover").Preload("Resources").Preload("Resources.UserRole").First(&performance, id).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &performance, nil
}

func (s *PerformanceStore) Delete(ctx context.Context, performance *model.Performance) error {
	for _, resource := range performance.Resources {
		err := os.Remove("resources/" + resource.Filename)
		if err != nil {
			return err
		}
	}
	err := s.db.WithContext(ctx).
		Delete(performance).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (s *PerformanceStore) SetCover(ctx context.Context, performance *model.Performance, resource *model.Resource) error {
	performance.Cover = resource
	if err := s.db.WithContext(ctx).Save(performance).Error; err != nil {
		return err
	}
	return nil
}

func (s *PerformanceStore) CheckPerformancesExist(ctx context.Context, performanceIds []uuid.UUID) ([]uuid.UUID, error) {
	var missing []uuid.UUID
	var uuids string
	for i, pid := range performanceIds {
		uuids += fmt.Sprintf("uuid('%s')", pid)
		if i < len(performanceIds)-1 {
			uuids += ","
		}
	}
	err := s.db.Raw(`
		SELECT x.id
		FROM unnest(ARRAY[` + uuids + `]::uuid[]) AS x(id)
		LEFT JOIN performances p ON p.id = x.id
		WHERE p.id IS NULL
	`).Scan(&missing).Error

	return missing, err
}
