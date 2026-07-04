package db

import (
	"context"
	"fmt"
	"my-gin-app/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func Connect(ctx context.Context, cfg config.Config) (*Mongo, error) {
	connectCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	clientOpts := options.Client().ApplyURI(cfg.MongoUri)
	client, err := mongo.Connect(connectCtx, clientOpts)
	if err != nil {
		return nil, fmt.Errorf("mongo connection failed %w", err)
	}
	database := client.Database(cfg.MongoDbName)
	fmt.Println("fromm db ", client, database)
	return &Mongo{
		Client: client,
		DB:     database,
	}, nil
}
