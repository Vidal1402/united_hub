package repository

import (
	"context"
	"time"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoNotificacaoRepository struct {
	coll *mongo.Collection
}

func NewMongoNotificacaoRepository(db *mongo.Database) *MongoNotificacaoRepository {
	return &MongoNotificacaoRepository{
		coll: db.Collection("notificacoes"),
	}
}

func (r *MongoNotificacaoRepository) ListEnviadas(ctx context.Context, pag PageParams) ([]domain.Notificacao, int64, error) {
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "criado_em", Value: -1}})

	cur, err := r.coll.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Notificacao
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.coll.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoNotificacaoRepository) Enviar(ctx context.Context, n *domain.Notificacao) error {
	now := time.Now().UTC()
	if n.CriadoEm.IsZero() {
		n.CriadoEm = now
	}
	_, err := r.coll.InsertOne(ctx, n)
	return err
}

