package store

import (
	"bandMate7/internal/model"
	"context"
	"mime/multipart"

	"github.com/google/uuid"
)

func NewMockStore() Storage {

	return Storage{
		Performances: &MockPerformancesStore{},
		Resources:    &MockResourcesStore{},
		UserRoles:    &MockUserRolesStore{},
		Artists:      &MockArtistsStore{},
		SetLists:     &MockSetlistsStore{},
	}
}

type MockPerformancesStore struct{}
type MockResourcesStore struct{}
type MockUserRolesStore struct{}
type MockArtistsStore struct{}
type MockSetlistsStore struct{}

func (mp *MockPerformancesStore) Create(ctx context.Context, performance *model.Performance) error {
	return nil
}
func (mp *MockPerformancesStore) GetAll(ctx context.Context, pq PaginatedQuery) ([]model.Performance, error) {
	return []model.Performance{}, nil
}
func (mp *MockPerformancesStore) Delete(ctx context.Context, performance *model.Performance) error {
	return nil
}
func (mp *MockPerformancesStore) GetByID(ctx context.Context, id uuid.UUID) (*model.Performance, error) {
	return &model.Performance{BaseModel: model.BaseModel{ID: id}}, nil
}
func (mp *MockPerformancesStore) SetCover(ctx context.Context, performance *model.Performance, resource *model.Resource) error {
	return nil
}
func (mp *MockPerformancesStore) CheckPerformancesExist(ctx context.Context, performanceIds []uuid.UUID) ([]uuid.UUID, error) {
	return []uuid.UUID{}, nil
}
func (mr *MockResourcesStore) Create(ctx context.Context, file multipart.File, header *multipart.FileHeader, performance *model.Performance, role *model.UserRole) (*model.Resource, error) {
	return &model.Resource{}, nil
}
func (mr *MockResourcesStore) Delete(ctx context.Context, resource model.Resource) error {
	return nil
}
func (mr *MockResourcesStore) GetByID(ctx context.Context, id uuid.UUID) (*model.Resource, error) {
	return &model.Resource{BaseModel: model.BaseModel{ID: id}}, nil
}
func (mr *MockResourcesStore) Open(ctx context.Context, resource *model.Resource) (*ResourceObject, error) {
	return &ResourceObject{}, nil
}
func (mr *MockResourcesStore) GetAllByPerformance(ctx context.Context, performance *model.Performance) ([]model.Resource, error) {
	return []model.Resource{}, nil
}
func (mur *MockUserRolesStore) GetByID(ctx context.Context, id uuid.UUID) (*model.UserRole, error) {
	return &model.UserRole{BaseModel: model.BaseModel{ID: id}}, nil
}
func (mur *MockUserRolesStore) GetAll(ctx context.Context) ([]model.UserRole, error) {
	return []model.UserRole{}, nil
}
func (ma *MockArtistsStore) GetAll(ctx context.Context) ([]model.Artist, error) {
	return []model.Artist{}, nil
}
func (ma *MockArtistsStore) Create(ctx context.Context, artist *model.Artist) error {
	return nil
}
func (ms *MockSetlistsStore) Create(ctx context.Context, setlist *model.Setlist, performanceIds []uuid.UUID) error {
	return nil
}
func (ms *MockSetlistsStore) GetAll(ctx context.Context) ([]model.Setlist, error) {
	return []model.Setlist{}, nil
}
func (ms *MockSetlistsStore) GetByID(ctx context.Context, id uuid.UUID, opts ...QueryOption) (*model.Setlist, error) {
	return &model.Setlist{BaseModel: model.BaseModel{ID: id}}, nil
}
func (ms *MockSetlistsStore) CheckPerformancesExist(ctx context.Context, setlist *model.Setlist, performanceIds []uuid.UUID) (missingPerfIds []uuid.UUID, notIncludedPerfIds []uuid.UUID, err error) {
	return []uuid.UUID{}, []uuid.UUID{}, nil
}
func (ms *MockSetlistsStore) UpdateOrder(ctx context.Context, setlist *model.Setlist, newOrder map[uuid.UUID]int) error {
	return nil
}
