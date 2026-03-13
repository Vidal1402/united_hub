package repository

import (
	"context"
	"time"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoAlertaRepository struct {
	coll *mongo.Collection
}

func NewMongoAlertaRepository(db *mongo.Database) *MongoAlertaRepository {
	return &MongoAlertaRepository{
		coll: db.Collection("alertas"),
	}
}

func (r *MongoAlertaRepository) List(ctx context.Context, pag PageParams) ([]domain.Alerta, int64, error) {
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cur, err := r.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Alerta
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoAlertaRepository) Resolver(ctx context.Context, uuid string) error {
	now := time.Now().UTC()
	update := bson.M{
		"$set": bson.M{
			"status":      "resolvido",
			"resolved_at": now,
		},
	}
	_, err := r.coll.UpdateOne(ctx, bson.M{"uuid": uuid}, update)
	return err
}

