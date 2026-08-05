package db

import (
	"bandMate7/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func New(debug bool) (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=band-organizer port=5432 sslmode=disable TimeZone=Europe/Berlin"

	logLevel := logger.Error

	if debug {
		logLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		return nil, err
	}

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	db.Exec(`
		DO $$
		BEGIN
		  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'resource_type') THEN
			CREATE TYPE resource_type AS ENUM ('image', 'video', 'audio', 'document');
		  END IF;
		END
		$$;
	`)

	db.SetupJoinTable(&model.Setlist{}, "Performances", &model.SetlistPerformance{})
	db.SetupJoinTable(&model.Performance{}, "Setlists", &model.SetlistPerformance{})

	err = db.AutoMigrate(
		&model.UserRole{},
		&model.Performance{},
		&model.Resource{},
		&model.Setlist{},
		&model.Artist{},
		&model.Meter{},
		&model.Collection{},
	)

	var userRoles = []model.UserRole{
		{Name: "Keyboard"},
		{Name: "Singer"},
		{Name: "Drums"},
		{Name: "Bass"},
		{Name: "Guitar 1"},
		{Name: "Guitar 2"},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&userRoles)

	if err != nil {
		return nil, err
	}

	return db, nil
}
