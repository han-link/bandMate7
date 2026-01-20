package store

import (
	"bandMate7/internal/model"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResourceStore struct {
	db *gorm.DB
}

func pathExists(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		// path/to/whatever exists
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		return false, nil
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		return true, err
	}
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
	var path string
	fileParts := strings.Split(resource.Filename, ".")
	name := fileParts[0]
	ext := fileParts[1]
	filename := resource.Filename
	version := 0
	for {
		path = "resources/" + filename
		version++
		exists, err := pathExists(path)
		if err != nil {
			return err, nil
		}
		if !exists {
			resource.Filename = filename
			break
		}
		filename = name + "_" + strconv.Itoa(version) + "." + ext
	}

	// TODO: Check if the file exists in the dir but not in he db
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

	dst, err := os.Create(path)
	if err != nil {
		return err, nil
	}

	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			// TODO: Handle Error
		}
	}(dst)

	if _, err := io.Copy(dst, file); err != nil {
		return err, nil
	}

	if err = file.Close(); err != nil {
		return err, nil
	}

	return nil, resource
}

func (s *ResourceStore) Delete(ctx context.Context, resource model.Resource) error {
	err := os.Remove("resources/" + resource.Filename)
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
