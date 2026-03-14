package repository

import (
	"context"
	"time"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoProdutoRepository struct {
	coll *mongo.Collection
}

func NewMongoProdutoRepository(db *mongo.Database) *MongoProdutoRepository {
	return &MongoProdutoRepository{
		coll: db.Collection("produtos"),
	}
}

func (r *MongoProdutoRepository) GetByUUID(ctx context.Context, uuid string) (*domain.Produto, error) {
	var p domain.Produto
	err := r.coll.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&p)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *MongoProdutoRepository) ListByFamilia(ctx context.Context, familia string, pag PageParams) ([]domain.Produto, int64, error) {
	filter := bson.M{"familia": familia}
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cur, err := r.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Produto
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoProdutoRepository) Create(ctx context.Context, p *domain.Produto) error {
	now := time.Now().UTC()
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
	p.UpdatedAt = now

	_, err := r.coll.InsertOne(ctx, p)
	return err
}

func (r *MongoProdutoRepository) Update(ctx context.Context, p *domain.Produto) error {
	p.UpdatedAt = time.Now().UTC()

	update := bson.M{
		"$set": bson.M{
			"familia":        p.Familia,
			"slug":           p.Slug,
			"nome":           p.Nome,
			"preco_centavos": p.Preco,
			"badge":          p.Badge,
			"descricao":      p.Descricao,
			"features":       p.Features,
			"updated_at":     p.UpdatedAt,
		},
	}
	_, err := r.coll.UpdateOne(ctx, bson.M{"uuid": p.UUID}, update)
	return err
}

func (r *MongoProdutoRepository) Delete(ctx context.Context, uuid string) error {
	_, err := r.coll.DeleteOne(ctx, bson.M{"uuid": uuid})
	return err
}

