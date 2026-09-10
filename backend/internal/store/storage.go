package store

import (
	"bandMate7/internal/model"
	"context"
	"errors"
	"mime/multipart"

	garage "git.deuxfleurs.fr/garage-sdk/garage-admin-sdk-golang"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

var (
	ErrNotFound         = errors.New("resource not found")
	ErrUserRoleNotFound = errors.New("user role not found")
	/*ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5*/
)

type Performances interface {
	Create(ctx context.Context, performance *model.Performance) error
	GetAll(ctx context.Context, pq PaginatedQuery) ([]model.Performance, error)
	Delete(ctx context.Context, performance *model.Performance) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Performance, error)
	SetCover(ctx context.Context, performance *model.Performance, resource *model.Resource) error
	CheckPerformancesExist(ctx context.Context, performanceIds []uuid.UUID) ([]uuid.UUID, error)
}

type Resources interface {
	Create(ctx context.Context, file multipart.File, header *multipart.FileHeader, performance *model.Performance, role *model.UserRole) (*model.Resource, error)
	Delete(ctx context.Context, resource model.Resource) error
	GetByID(ctx context.Context, id uuid.UUID) (*model.Resource, error)
	Open(ctx context.Context, resource *model.Resource) (*ResourceObject, error)
	GetAllByPerformance(ctx context.Context, performance *model.Performance) ([]model.Resource, error)
}

type UserRoles interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.UserRole, error)
	GetAll(ctx context.Context) ([]model.UserRole, error)
}

type Artists interface {
	GetAll(ctx context.Context) ([]model.Artist, error)
	Create(ctx context.Context, artist *model.Artist) error
}

type SetLists interface {
	Create(ctx context.Context, setlist *model.Setlist, performanceIds []uuid.UUID) error
	GetAll(ctx context.Context) ([]model.Setlist, error)
	GetByID(ctx context.Context, id uuid.UUID, opts ...QueryOption) (*model.Setlist, error)
	CheckPerformancesExist(ctx context.Context, setlist *model.Setlist, performanceIds []uuid.UUID) (missingPerfIds []uuid.UUID, notIncludedPerfIds []uuid.UUID, err error)
	UpdateOrder(ctx context.Context, setlist *model.Setlist, newOrder map[uuid.UUID]int) error
}

type Storage struct {
	Performances Performances
	Resources    Resources
	UserRoles    UserRoles
	Artists      Artists
	SetLists     SetLists
}

func NewStorage(db *gorm.DB, resourceDir string, garageClient *garage.APIClient, garageCtx context.Context, minioClient *minio.Client, bucket string) Storage {
	return Storage{
		Performances: &PerformanceStore{db},
		Resources:    &ResourceStore{db, resourceDir, garageClient, garageCtx, minioClient, bucket},
		UserRoles:    &UserRoleStore{db},
		Artists:      &ArtistStore{db},
		SetLists:     &SetlistStore{db},
	}
}
