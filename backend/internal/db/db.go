package db

import (
	"bandMate7/internal/model"
	"fmt"
	"log"
	_ "time/tzdata" // Fallback if timezone database is not present on host

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func New(debug bool, host string, user string, password string, dbName string, port int, sslMode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC", host, user, password, dbName, port, sslMode)
	logLevel := logger.Error

	if debug {
		logLevel = logger.Info
		log.Print(dsn)
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

	err = db.AutoMigrate(
		&model.UserRole{},
		&model.Performance{},
		&model.Resource{},
		&model.Setlist{},
		&model.Artist{},
		&model.Meter{},
		&model.Collection{},
		&model.SetlistPerformance{},
	)

	var userRoles = []model.UserRole{
		{Titel: "Keyboard"},
		{Titel: "Singer"},
		{Titel: "Drums"},
		{Titel: "Bass"},
		{Titel: "Guitar 1"},
		{Titel: "Guitar 2"},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&userRoles)

	var meters = []model.Meter{
		{Titel: "3/4"},
		{Titel: "4/4"},
		{Titel: "6/8"},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&meters)

	var genres = []model.Genre{
		{Titel: "Rock"},
		{Titel: "Pop"},
	}
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&genres)

	if err != nil {
		return nil, err
	}

	return db, nil
}
