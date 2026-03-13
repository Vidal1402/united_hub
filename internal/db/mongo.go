package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongo(ctx context.Context, uri, dbName string) (*Mongo, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri).
		SetConnectTimeout(10 * time.Second).
		SetServerSelectionTimeout(10 * time.Second)

	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(ctx)
		return nil, err
	}

	return &Mongo{
		Client:   client,
		Database: client.Database(dbName),
	}, nil
}

func (m *Mongo) Close(ctx context.Context) error {
	if m == nil || m.Client == nil {
		return nil
	}
	return m.Client.Disconnect(ctx)
}
