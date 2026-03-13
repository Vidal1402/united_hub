package repository

import (
	"context"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoKanbanRepository struct {
	collColumns *mongo.Collection
	collCards   *mongo.Collection
}

func NewMongoKanbanRepository(db *mongo.Database) *MongoKanbanRepository {
	return &MongoKanbanRepository{
		collColumns: db.Collection("kanban_columns"),
		collCards:   db.Collection("kanban_cards"),
	}
}

func (r *MongoKanbanRepository) ListColumnsByCliente(ctx context.Context, clienteUUID string) ([]domain.KanbanColumn, error) {
	filter := bson.M{"cliente_uuid": clienteUUID}
	opts := options.Find().SetSort(bson.D{{Key: "ordenacao", Value: 1}})

	cur, err := r.collColumns.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var out []domain.KanbanColumn
	if err := cur.All(ctx, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (r *MongoKanbanRepository) ListCardsByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.KanbanCard, int64, error) {
	filter := bson.M{"cliente_uuid": clienteUUID}
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "prazo", Value: 1}})

	cur, err := r.collCards.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.KanbanCard
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.collCards.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

