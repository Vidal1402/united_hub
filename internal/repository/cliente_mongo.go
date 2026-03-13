package repository

import (
	"context"
	"time"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClienteRepository struct {
	coll *mongo.Collection
}

func NewMongoClienteRepository(db *mongo.Database) *MongoClienteRepository {
	return &MongoClienteRepository{
		coll: db.Collection("clientes"),
	}
}

func (r *MongoClienteRepository) GetByUUID(ctx context.Context, uuid string) (*domain.Cliente, error) {
	var c domain.Cliente
	err := r.coll.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&c)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *MongoClienteRepository) List(ctx context.Context, pag PageParams) ([]domain.Cliente, int64, error) {
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cur, err := r.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Cliente
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoClienteRepository) ListByOwner(ctx context.Context, ownerUUID string, pag PageParams) ([]domain.Cliente, int64, error) {
	filter := bson.M{"owner_uuid": ownerUUID}

	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cur, err := r.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Cliente
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoClienteRepository) Create(ctx context.Context, c *domain.Cliente) error {
	now := time.Now().UTC()
	if c.CreatedAt.IsZero() {
		c.CreatedAt = now
	}
	c.UpdatedAt = now

	_, err := r.coll.InsertOne(ctx, c)
	return err
}

func (r *MongoClienteRepository) Update(ctx context.Context, c *domain.Cliente) error {
	c.UpdatedAt = time.Now().UTC()

	update := bson.M{
		"$set": bson.M{
			"nome":       c.Nome,
			"email":      c.Email,
			"segmento":   c.Segmento,
			"plano":      c.Plano,
			"status":     c.Status,
			"cidade":     c.Cidade,
			"owner_uuid": c.OwnerUUID,
			"updated_at": c.UpdatedAt,
		},
	}

	_, err := r.coll.UpdateOne(ctx, bson.M{"uuid": c.UUID}, update)
	return err
}

func (r *MongoClienteRepository) Desativar(ctx context.Context, uuid string) error {
	update := bson.M{
		"$set": bson.M{
			"status":     "inativo",
			"updated_at": time.Now().UTC(),
		},
	}
	_, err := r.coll.UpdateOne(ctx, bson.M{"uuid": uuid}, update)
	return err
}

