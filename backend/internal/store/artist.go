package store

import (
	"bandMate7/internal/model"
	"context"

	"gorm.io/gorm"
)

type ArtistStore struct {
	db *gorm.DB
}

func (s *ArtistStore) GetAll(ctx context.Context) ([]model.Artist, error) {
	var artists []model.Artist
	err := s.db.WithContext(ctx).Find(&artists).Error
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func (s *ArtistStore) Create(ctx context.Context, artist *model.Artist) error {
	err := s.db.WithContext(ctx).
		Create(artist).
		Error
	if err != nil {
		return err
	}
	return nil
}
