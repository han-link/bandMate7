package store

import (
	"bandMate7/internal/model"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	garage "git.deuxfleurs.fr/garage-sdk/garage-admin-sdk-golang"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type ResourceStore struct {
	db           *gorm.DB
	resourceDir  string
	garageClient *garage.APIClient
	garageCtx    context.Context
	minioClient  *minio.Client
	bucket       string
}

func (s *ResourceStore) objectExists(ctx context.Context, key string) (bool, error) {
	_, err := s.minioClient.StatObject(ctx, s.bucket, key, minio.StatObjectOptions{})
	if err != nil {
		if minio.ToErrorResponse(err).StatusCode == http.StatusNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *ResourceStore) Create(
	ctx context.Context,
	file multipart.File,
	header *multipart.FileHeader,
	performance *model.Performance,
	role *model.UserRole,
) (error, *model.Resource) {
	resource := &model.Resource{
		Type:          model.ResourceTypeImage, // TODO: Replace with type detection
		Filename:      header.Filename,
		SizeBytes:     header.Size,
		PerformanceID: performance.ID,
		UserRole:      role,
	}

	name, ext, _ := strings.Cut(resource.Filename, ".")
	if ext != "" {
		ext = "." + ext
	}
	objectKey := resource.Filename
	for version := 1; ; version++ {
		exists, err := s.objectExists(ctx, objectKey)
		if err != nil {
			return err, nil
		}
		if !exists {
			break
		}
		objectKey = fmt.Sprintf("%s_%d%s", name, version, ext)
	}
	resource.Filename = objectKey

	err := s.db.WithContext(ctx).
		Create(resource).
		Error
	if err != nil {
		return err, nil
	}
	err = s.db.WithContext(ctx).
		Preload("Resources").
		First(performance).
		Error
	if err != nil {
		return err, nil
	}

	contentType := header.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	if _, err := s.minioClient.PutObject(ctx, s.bucket, objectKey, file, header.Size, minio.PutObjectOptions{
		ContentType: contentType,
	}); err != nil {
		return err, nil
	}

	if err := file.Close(); err != nil {
		return err, nil
	}

	return nil, resource
}

func (s *ResourceStore) Delete(ctx context.Context, resource model.Resource) error {
	err := os.Remove(path.Join(s.resourceDir, resource.Filename))
	if err != nil {
		return err
	}
	err = s.db.WithContext(ctx).
		Delete(&resource).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (s *ResourceStore) GetByID(ctx context.Context, id uuid.UUID) (*model.Resource, error) {
	var resource model.Resource
	err := s.db.WithContext(ctx).First(&resource, id).Error
	if err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &resource, nil
}

type ResourceObject struct {
	io.ReadSeekCloser
	ContentType  string
	Size         int64
	LastModified time.Time
}

func (s *ResourceStore) Open(ctx context.Context, resource *model.Resource) (*ResourceObject, error) {
	obj, err := s.minioClient.GetObject(ctx, s.bucket, resource.Filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	stat, err := obj.Stat()
	if err != nil {
		obj.Close()
		if minio.ToErrorResponse(err).StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &ResourceObject{
		ReadSeekCloser: obj,
		ContentType:    stat.ContentType,
		Size:           stat.Size,
		LastModified:   stat.LastModified,
	}, nil
}

func (s *ResourceStore) GetAllByPerformance(ctx context.Context, performance *model.Performance) ([]model.Resource, error) {
	var resources []model.Resource
	err := s.db.WithContext(ctx).
		Where(&model.Resource{PerformanceID: performance.ID}).
		Find(&resources).
		Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}
