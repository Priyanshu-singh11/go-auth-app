package app

import (
	"context"
	"fmt"
	"my-gin-app/internal/config"
	"my-gin-app/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Config      config.Config
	MongoClient *mongo.Client
	DB          *mongo.Database
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	mongoCli, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, err
	}
	fmt.Println(cfg, mongoCli)
	return &App{
		Config:      cfg,
		MongoClient: mongoCli.Client,
		DB:          mongoCli.DB,
	}, nil
}

func (a *App) Close(ctx context.Context) error {
	if a.MongoClient == nil {
		return nil
	}
	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("mongo disconnect failed %w", err)
	}
	return nil
}
