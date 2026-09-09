package store

import (
	"bandMate7/internal/model"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SetlistStore struct {
	db *gorm.DB
}

func (s *SetlistStore) Create(ctx context.Context, setlist *model.Setlist, performanceIds []uuid.UUID) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Omit("Performances").Create(&setlist).Error; err != nil {
			return err
		}

		joins := make([]model.SetlistPerformance, len(performanceIds))
		for i, pid := range performanceIds {
			joins[i] = model.SetlistPerformance{
				SetlistID:     setlist.ID,
				PerformanceID: pid,
				Position:      i + 1,
			}
		}
		return tx.Create(joins).Error
	})

	if err != nil {
		return err
	}

	if err := s.db.WithContext(ctx).Scopes(WithSortedSetlist()).Take(&setlist).Error; err != nil {
		return err
	}

	return nil
}

func (s *SetlistStore) GetAll(ctx context.Context) ([]model.Setlist, error) {
	var setlists []model.Setlist
	err := s.db.WithContext(ctx).Find(&setlists).Error
	if err != nil {
		return nil, err
	}
	return setlists, nil
}

func (s *SetlistStore) GetByID(ctx context.Context, id uuid.UUID, opts ...QueryOption) (*model.Setlist, error) {
	var setlist model.Setlist
	err := s.db.WithContext(ctx).Scopes(opts...).First(&setlist, id).Error

	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &setlist, nil
}

func (s *SetlistStore) CheckPerformancesExist(ctx context.Context, setlist *model.Setlist, performanceIds []uuid.UUID) (missingPerfIds []uuid.UUID, notIncludedPerfIds []uuid.UUID, err error) {
	var missing []uuid.UUID
	var notIncluded []uuid.UUID
	var uuids string
	for i, pid := range performanceIds {
		uuids += fmt.Sprintf("uuid('%s')", pid)
		if i < len(performanceIds)-1 {
			uuids += ","
		}
	}
	err = s.db.Raw(`
		SELECT sp.performance_id
		FROM unnest(ARRAY[` + uuids + `]::uuid[]) AS x(id)
		RIGHT JOIN setlist_performances sp ON sp.performance_id = x.id
        WHERE x.id IS NULL
        AND sp.setlist_id = uuid('` + setlist.ID.String() + `')
	`).Scan(&missing).Error

	if err != nil {
		return missing, notIncluded, err
	}

	err = s.db.Raw(`
		SELECT x.id
		FROM unnest(ARRAY[` + uuids + `]::uuid[]) AS x(id)
		WHERE NOT EXISTS (
			SELECT * FROM setlist_performances sp WHERE sp.performance_id = x.id
			AND sp.setlist_id = uuid('` + setlist.ID.String() + `')
		)
	`).Scan(&notIncluded).Error

	return missing, notIncluded, err
}

func (s *SetlistStore) UpdateOrder(ctx context.Context, setlist *model.Setlist, newOrder map[uuid.UUID]int) error {
	if len(setlist.Items) != len(newOrder) {
		return fmt.Errorf("expected %d positions, got %d", len(setlist.Items), len(newOrder))
	}

	for i := range setlist.Items {
		newPos := newOrder[setlist.Items[i].PerformanceID]
		setlist.Items[i].Position = newPos
	}

	if err := s.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&setlist).Error; err != nil {
		return err
	}

	return nil
}
