package store

import (
	"bandMate7/internal/model"
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRoleStore struct {
	db *gorm.DB
}

func (s *UserRoleStore) GetByID(ctx context.Context, id uuid.UUID) (*model.UserRole, error) {
	var role model.UserRole
	err := s.db.WithContext(ctx).First(&role, id).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &role, nil
}

func (s *UserRoleStore) GetAll(ctx context.Context) ([]model.UserRole, error) {
	var roles []model.UserRole
	err := s.db.WithContext(ctx).Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}
