package main

import (
	"bandMate7/internal/db"
	"bandMate7/internal/store"
	"encoding/json"
	"log"
	"os"

	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			BandMate7 API
//	@description	API for bandMate7, an expense manager

// @description
func main() {
	cfg := config{
		addr:   ":8080",
		domain: "localhost",
	}

	logDir := "./logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("failed to create log directory: %v", err)
	}

	// Logger
	loggerConfig := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "./logs/app.log"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	// Unmarshal the JSON configuration into a zap.Config
	var loggerCfg zap.Config
	if err := json.Unmarshal(loggerConfig, &loggerCfg); err != nil {
		log.Fatalf("Error unmarshaling zap config: %v", err)
	}

	// Build the logger from the custom configuration
	must := zap.Must(loggerCfg.Build())
	defer must.Sync()

	// Create a sugared logger from the built logger
	logger := must.Sugar()

	logger.Info("Sugared logger constructed successfully")

	database, err := db.New(false)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Database connection established")

	storage := store.NewStorage(database)

	app := &application{
		config: cfg,
		logger: logger,
		store:  storage,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
