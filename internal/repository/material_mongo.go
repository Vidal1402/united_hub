package repository

import (
	"context"
	"time"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoMaterialRepository struct {
	collPastas  *mongo.Collection
	collArquivos *mongo.Collection
}

func NewMongoMaterialRepository(db *mongo.Database) *MongoMaterialRepository {
	return &MongoMaterialRepository{
		collPastas:  db.Collection("materiais_pastas"),
		collArquivos: db.Collection("materiais_arquivos"),
	}
}

func (r *MongoMaterialRepository) ListPastasByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.MaterialPasta, int64, error) {
	filter := bson.M{"cliente_uuid": clienteUUID}
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cur, err := r.collPastas.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.MaterialPasta
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.collPastas.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoMaterialRepository) ListArquivosByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.MaterialArquivo, int64, error) {
	filter := bson.M{"cliente_uuid": clienteUUID}
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "data", Value: -1}})

	cur, err := r.collArquivos.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.MaterialArquivo
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.collArquivos.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoMaterialRepository) CreatePasta(ctx context.Context, p *domain.MaterialPasta) error {
	now := time.Now().UTC()
	if p.CreatedAt.IsZero() {
		p.CreatedAt = now
	}
	p.UpdatedAt = now
	_, err := r.collPastas.InsertOne(ctx, p)
	return err
}

func (r *MongoMaterialRepository) CreateArquivo(ctx context.Context, a *domain.MaterialArquivo) error {
	now := time.Now().UTC()
	if a.CreatedAt.IsZero() {
		a.CreatedAt = now
	}
	a.UpdatedAt = now

	_, err := r.collArquivos.InsertOne(ctx, a)
	return err
}

func (r *MongoMaterialRepository) GetPastaByUUID(ctx context.Context, uuid string) (*domain.MaterialPasta, error) {
	var p domain.MaterialPasta
	err := r.collPastas.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&p)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *MongoMaterialRepository) UpdatePasta(ctx context.Context, p *domain.MaterialPasta) error {
	p.UpdatedAt = time.Now().UTC()
	_, err := r.collPastas.ReplaceOne(ctx, bson.M{"uuid": p.UUID}, p)
	return err
}

func (r *MongoMaterialRepository) GetArquivoByUUID(ctx context.Context, uuid string) (*domain.MaterialArquivo, error) {
	var a domain.MaterialArquivo
	err := r.collArquivos.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&a)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (r *MongoMaterialRepository) UpdateArquivo(ctx context.Context, a *domain.MaterialArquivo) error {
	a.UpdatedAt = time.Now().UTC()
	_, err := r.collArquivos.ReplaceOne(ctx, bson.M{"uuid": a.UUID}, a)
	return err
}

