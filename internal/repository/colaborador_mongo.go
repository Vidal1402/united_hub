package repository

import (
	"context"
	"time"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoColaboradorRepository struct {
	coll *mongo.Collection
}

func NewMongoColaboradorRepository(db *mongo.Database) *MongoColaboradorRepository {
	return &MongoColaboradorRepository{
		coll: db.Collection("colaboradores"),
	}
}

func (r *MongoColaboradorRepository) GetByUUID(ctx context.Context, uuid string) (*domain.Colaborador, error) {
	var c domain.Colaborador
	err := r.coll.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&c)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *MongoColaboradorRepository) List(ctx context.Context, pag PageParams) ([]domain.Colaborador, int64, error) {
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cur, err := r.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Colaborador
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoColaboradorRepository) Create(ctx context.Context, c *domain.Colaborador) error {
	now := time.Now().UTC()
	if c.CreatedAt.IsZero() {
		c.CreatedAt = now
	}
	c.UpdatedAt = now

	_, err := r.coll.InsertOne(ctx, c)
	return err
}

func (r *MongoColaboradorRepository) Update(ctx context.Context, c *domain.Colaborador) error {
	c.UpdatedAt = time.Now().UTC()

	update := bson.M{
		"$set": bson.M{
			"nome":       c.Nome,
			"email":      c.Email,
			"cargo":      c.Cargo,
			"role":       c.Role,
			"status":     c.Status,
			"updated_at": c.UpdatedAt,
		},
	}

	_, err := r.coll.UpdateOne(ctx, bson.M{"uuid": c.UUID}, update)
	return err
}

