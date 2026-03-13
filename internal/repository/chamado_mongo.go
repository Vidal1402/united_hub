package repository

import (
	"context"
	"time"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoChamadoRepository struct {
	coll *mongo.Collection
}

func NewMongoChamadoRepository(db *mongo.Database) *MongoChamadoRepository {
	return &MongoChamadoRepository{
		coll: db.Collection("chamados"),
	}
}

func (r *MongoChamadoRepository) ListByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.Chamado, int64, error) {
	filter := bson.M{"cliente_uuid": clienteUUID}
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "criado_em", Value: -1}})

	cur, err := r.coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Chamado
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoChamadoRepository) Create(ctx context.Context, c *domain.Chamado) error {
	now := time.Now().UTC()
	if c.CriadoEm.IsZero() {
		c.CriadoEm = now
	}
	c.AtualizadoEm = now

	_, err := r.coll.InsertOne(ctx, c)
	return err
}

