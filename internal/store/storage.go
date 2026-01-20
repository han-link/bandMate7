package store

import (
	"bandMate7/internal/model"
	"context"
	"errors"
	"mime/multipart"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("resource not found")
	/*ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5*/
)

type Performances interface {
	Create(ctx context.Context, performance *model.Performance) error
	GetAll(ctx context.Context) ([]model.Performance, error)
	Delete(ctx context.Context, performance *model.Performance) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Performance, error)
	SetCover(ctx context.Context, performance *model.Performance, resource *model.Resource) error
}

type Resources interface {
	Create(ctx context.Context, file multipart.File, header *multipart.FileHeader, performance *model.Performance, role *model.UserRole) (error, *model.Resource)
	Delete(ctx context.Context, resource model.Resource) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Resource, error)
	GetAllByPerformance(ctx context.Context, performance *model.Performance) ([]model.Resource, error)
}

type UserRoles interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.UserRole, error)
	GetAll(ctx context.Context) ([]model.UserRole, error)
}

type Storage struct {
	Performances Performances
	Resources    Resources
	UserRoles    UserRoles
}

func NewStorage(db *gorm.DB) Storage {
	return Storage{
		Performances: &PerformanceStore{db},
		Resources:    &ResourceStore{db},
		UserRoles:    &UserRoleStore{db},
	}
}
