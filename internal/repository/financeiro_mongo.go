package repository

import (
	"context"
	"time"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoFinanceiroRepository struct {
	collFaturas    *mongo.Collection
	collRecebiveis *mongo.Collection
	collPagamentos *mongo.Collection
}

func NewMongoFinanceiroRepository(db *mongo.Database) *MongoFinanceiroRepository {
	return &MongoFinanceiroRepository{
		collFaturas:    db.Collection("faturas"),
		collRecebiveis: db.Collection("recebiveis"),
		collPagamentos: db.Collection("pagamentos"),
	}
}

func (r *MongoFinanceiroRepository) ListFaturasByCliente(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.Fatura, int64, error) {
	filter := bson.M{"cliente_uuid": clienteUUID}
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "vencimento", Value: -1}})

	cur, err := r.collFaturas.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Fatura
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.collFaturas.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoFinanceiroRepository) ListRecebiveis(ctx context.Context, pag PageParams) ([]domain.Recebivel, int64, error) {
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "vencimento", Value: -1}})

	cur, err := r.collRecebiveis.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Recebivel
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.collRecebiveis.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoFinanceiroRepository) ListPagamentos(ctx context.Context, pag PageParams) ([]domain.Pagamento, int64, error) {
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "vencimento", Value: -1}})

	cur, err := r.collPagamentos.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Pagamento
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.collPagamentos.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoFinanceiroRepository) CreateRecebivel(ctx context.Context, rec *domain.Recebivel) error {
	now := time.Now().UTC()
	if rec.CreatedAt.IsZero() {
		rec.CreatedAt = now
	}
	rec.UpdatedAt = now

	_, err := r.collRecebiveis.InsertOne(ctx, rec)
	return err
}

func (r *MongoFinanceiroRepository) MarkRecebivelPago(ctx context.Context, uuid string) error {
	now := time.Now().UTC()
	update := bson.M{
		"$set": bson.M{
			"status":     "pago",
			"updated_at": now,
		},
	}
	_, err := r.collRecebiveis.UpdateOne(ctx, bson.M{"uuid": uuid}, update)
	return err
}

