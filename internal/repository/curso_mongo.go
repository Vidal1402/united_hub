package repository

import (
	"context"

	"backend_united_hub/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCursoRepository struct {
	collCursos    *mongo.Collection
	collProgresso *mongo.Collection
}

func NewMongoCursoRepository(db *mongo.Database) *MongoCursoRepository {
	return &MongoCursoRepository{
		collCursos:    db.Collection("cursos"),
		collProgresso: db.Collection("cursos_progresso"),
	}
}

func (r *MongoCursoRepository) ListCursos(ctx context.Context, pag PageParams) ([]domain.Curso, int64, error) {
	opts := options.Find().
		SetLimit(int64(pag.Limit)).
		SetSkip(int64(pag.Offset)).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cur, err := r.collCursos.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cur.Close(ctx)

	var out []domain.Curso
	if err := cur.All(ctx, &out); err != nil {
		return nil, 0, err
	}

	total, err := r.collCursos.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return out, total, nil
}

func (r *MongoCursoRepository) ListCursosComProgresso(ctx context.Context, clienteUUID string, pag PageParams) ([]domain.Curso, []domain.CursoProgresso, int64, error) {
	// paginação em cursos, progresso filtrado por cliente
	cursos, total, err := r.ListCursos(ctx, pag)
	if err != nil {
		return nil, nil, 0, err
	}

	filter := bson.M{"cliente_uuid": clienteUUID}
	cur, err := r.collProgresso.Find(ctx, filter)
	if err != nil {
		return nil, nil, 0, err
	}
	defer cur.Close(ctx)

	var progresso []domain.CursoProgresso
	if err := cur.All(ctx, &progresso); err != nil {
		return nil, nil, 0, err
	}

	return cursos, progresso, total, nil
}

