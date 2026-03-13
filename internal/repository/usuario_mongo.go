package repository

import (
	"context"
	"time"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoUsuarioRepository struct {
	coll *mongo.Collection
}

func NewMongoUsuarioRepository(db *mongo.Database) *MongoUsuarioRepository {
	return &MongoUsuarioRepository{
		coll: db.Collection("usuarios"),
	}
}

func (r *MongoUsuarioRepository) GetByUUID(ctx context.Context, uuid string) (*domain.Usuario, error) {
	var u domain.Usuario
	err := r.coll.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *MongoUsuarioRepository) GetByEmail(ctx context.Context, email string) (*domain.Usuario, error) {
	var u domain.Usuario
	err := r.coll.FindOne(ctx, bson.M{"email": email}).Decode(&u)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *MongoUsuarioRepository) List(ctx context.Context, pag PageParams) ([]domain.Usuario, int64, error) {
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cur, err := r.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Usuario
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoUsuarioRepository) Create(ctx context.Context, u *domain.Usuario) error {
	now := time.Now().UTC()
	if u.CreatedAt.IsZero() {
		u.CreatedAt = now
	}
	u.UpdatedAt = now

	_, err := r.coll.InsertOne(ctx, u)
	return err
}

func (r *MongoUsuarioRepository) Update(ctx context.Context, u *domain.Usuario) error {
	u.UpdatedAt = time.Now().UTC()

	update := bson.M{
		"$set": bson.M{
			"cliente_uuid":     u.ClienteUUID,
			"email":            u.Email,
			"senha_hash":       u.SenhaHash,
			"role":             u.Role,
			"can_producao":     u.CanProducao,
			"can_performance":  u.CanPerformance,
			"updated_at":       u.UpdatedAt,
		},
	}
	_, err := r.coll.UpdateOne(ctx, bson.M{"uuid": u.UUID}, update)
	return err
}

