package store

import "gorm.io/gorm"

type QueryOption = func(*gorm.DB) *gorm.DB

func WithSortedSetlist() QueryOption {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Preload("Items", func(db *gorm.DB) *gorm.DB {
				return db.Order("setlist_performances.position ASC")
			}).
			Preload("Items.Performance")
	}
}
