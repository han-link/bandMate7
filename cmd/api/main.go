package main

import (
	"bandMate7/internal/db"
	"bandMate7/internal/env"
	"bandMate7/internal/service"
	"bandMate7/internal/store"
	"context"
	"encoding/json"
	"log"
	"os"

	garage "git.deuxfleurs.fr/garage-sdk/garage-admin-sdk-golang"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			BandMate7 API
//	@description	API for bandMate7, an expense manager

// @description
func main() {
	rootDir := env.GetString("Resource_DIR", "No path provided")
	if rootDir == "No path provided" {
		log.Fatal("No resource path provided. Set env variable Resource_DIR")
	}

	cfg := config{
		addr:        env.GetString("ADDR", ":8080"),
		host:        env.GetString("HOST", "localhost"),
		resourceDir: rootDir,
	}

	cfg.baseUrl = "http://" + cfg.host + cfg.addr

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

	garageCfg := garage.NewConfiguration()

	client := garage.NewAPIClient(garageCfg)

	ctx := context.WithValue(context.Background(), garage.ContextAccessToken, env.GetString("GARAGE_DEFAULT_ACCESS_KEY", ""))

	garageEndpoint := env.GetString("GARAGE_S3_ENDPOINT", "localhost:3900")
	garageAccessKey := env.GetString("GARAGE_DEFAULT_ACCESS_KEY", "")
	garageSecretKey := env.GetString("GARAGE_DEFAULT_SECRET_KEY", "")
	garageBucket := env.GetString("GARAGE_DEFAULT_BUCKET", "band-organizer")

	minioClient, err := minio.New(garageEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(garageAccessKey, garageSecretKey, ""),
		Secure: env.GetBool("GARAGE_S3_USE_SSL", false),
		Region: env.GetString("GARAGE_REGION", "garage"),
	})
	if err != nil {
		logger.Fatal(err)
	}

	storage := store.NewStorage(database, cfg.resourceDir, client, ctx, minioClient, garageBucket)

	services := service.NewServices(&storage, cfg.baseUrl)

	app := &application{
		config:  cfg,
		logger:  logger,
		store:   storage,
		service: services,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
