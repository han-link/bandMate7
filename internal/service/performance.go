package service

import (
	"bandMate7/internal/model"
	"bandMate7/internal/store"
	"context"
	"errors"
	"mime/multipart"

	"github.com/google/uuid"
)

var ErrUserRoleIdRequiredForScore = errors.New("userRoleId is required when appending a file")

type PerformanceService struct {
	baseUrl string
	store   *store.Storage
}

type CreatePerformanceRequest struct {
	Name        string
	Bpm         *int
	UserRoleId  *uuid.UUID
	Cover       multipart.File
	CoverHeader *multipart.FileHeader
	Score       multipart.File
	ScoreHeader *multipart.FileHeader
}

func (s *PerformanceService) Create(ctx context.Context, input CreatePerformanceRequest) (*model.Performance, error) {
	if input.Score != nil && input.UserRoleId == nil {
		return nil, ErrUserRoleIdRequiredForScore
	}

	performance := &model.Performance{Name: input.Name}
	if input.Bpm != nil {
		performance.Bpm = input.Bpm
	}

	if err := s.store.Performances.Create(ctx, performance); err != nil {
		return nil, err
	}

	if input.Cover != nil {
		err, resource := s.store.Resources.Create(ctx, input.Cover, input.CoverHeader, performance, nil)
		if err != nil {
			return nil, err
		}
		if err = s.store.Performances.SetCover(ctx, performance, resource); err != nil {
			return nil, err
		}
	}

	if input.UserRoleId != nil {
		userRole, err := s.store.UserRoles.GetByID(ctx, *input.UserRoleId)
		if err != nil {
			return nil, err
		}
		if input.Score != nil {
			err, _ = s.store.Resources.Create(ctx, input.Score, input.ScoreHeader, performance, userRole)
			if err != nil {
				return nil, err
			}
		}
	}
	return performance, nil
}

func (s *PerformanceService) performanceAddCoverUrl(p *model.Performance) {
	if p == nil {
		return
	}
	if p.Cover == nil {
		return
	}
	p.Cover.SetUrl(s.baseUrl)
}

func (s *PerformanceService) GetAll(ctx context.Context) (*[]model.Performance, error) {
	performances, err := s.store.Performances.GetAll(ctx)

	for i := range performances {
		s.performanceAddCoverUrl(&performances[i])
	}

	return &performances, err
}
