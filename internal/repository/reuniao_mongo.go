package repository

import (
	"context"
	"time"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoReuniaoRepository struct {
	coll *mongo.Collection
}

func NewMongoReuniaoRepository(db *mongo.Database) *MongoReuniaoRepository {
	return &MongoReuniaoRepository{
		coll: db.Collection("reunioes"),
	}
}

func (r *MongoReuniaoRepository) ListProximasByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.Reuniao, int64, error) {
	now := time.Now()
	filter := bson.M{
		"cliente_uuid": clienteUUID,
		"data_hora":    bson.M{"$gte": now},
	}
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "data_hora", Value: 1}})

	cur, err := r.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Reuniao
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoReuniaoRepository) ListHistoricoByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.Reuniao, int64, error) {
	now := time.Now()
	filter := bson.M{
		"cliente_uuid": clienteUUID,
		"data_hora":    bson.M{"$lt": now},
	}
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "data_hora", Value: -1}})

	cur, err := r.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Reuniao
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoReuniaoRepository) ListAdmin(ctx context.Context, pag PageParams) ([]domain.Reuniao, int64, error) {
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "data_hora", Value: -1}})

	cur, err := r.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Reuniao
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoReuniaoRepository) Create(ctx context.Context, reun *domain.Reuniao) error {
	now := time.Now().UTC()
	if reun.CreatedAt.IsZero() {
		reun.CreatedAt = now
	}
	reun.UpdatedAt = now
	_, err := r.coll.InsertOne(ctx, reun)
	return err
}

