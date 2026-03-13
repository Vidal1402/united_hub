package repository

import (
	"context"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRelatorioRepository struct {
	coll *mongo.Collection
}

func NewMongoRelatorioRepository(db *mongo.Database) *MongoRelatorioRepository {
	return &MongoRelatorioRepository{
		coll: db.Collection("relatorios"),
	}
}

func (r *MongoRelatorioRepository) ListByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.Relatorio, int64, error) {
	filter := bson.M{"cliente_uuid": clienteUUID}
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "data", Value: -1}})

	cur, err := r.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Relatorio
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return out, total, nil
}

func (r *MongoRelatorioRepository) ListAdmin(ctx context.Context, pag PageParams) ([]domain.Relatorio, int64, error) {
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "data", Value: -1}})

	cur, err := r.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Relatorio
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

