package migrations

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpMongo: garante collections + índices
func UpMongo(ctx context.Context, db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	// 1) Cliente
	if err := ensureIndexes(ctx, db.Collection("clientes"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 2) Colaborador
	if err := ensureIndexes(ctx, db.Collection("colaboradores"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 3) KanbanColumn
	if err := ensureIndexes(ctx, db.Collection("kanban_columns"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "cliente_uuid", Value: 1}}},
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 4) KanbanCard
	if err := ensureIndexes(ctx, db.Collection("kanban_cards"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "cliente_uuid", Value: 1}, {Key: "column_uuid", Value: 1}}},
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 5) Relatorio
	if err := ensureIndexes(ctx, db.Collection("relatorios"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "cliente_uuid", Value: 1}, {Key: "periodo", Value: 1}}},
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 6) MaterialPasta
	if err := ensureIndexes(ctx, db.Collection("materiais_pastas"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "cliente_uuid", Value: 1}}},
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 7) MaterialArquivo
	if err := ensureIndexes(ctx, db.Collection("materiais_arquivos"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "cliente_uuid", Value: 1}, {Key: "pasta_uuid", Value: 1}}},
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 8) Reuniao
	if err := ensureIndexes(ctx, db.Collection("reunioes"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "cliente_uuid", Value: 1}, {Key: "data_hora", Value: -1}}},
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 9) Fatura
	if err := ensureIndexes(ctx, db.Collection("faturas"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "cliente_uuid", Value: 1}, {Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "vencimento", Value: -1}}},
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 10) Curso
	if err := ensureIndexes(ctx, db.Collection("cursos"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "slug", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 11) CursoProgresso
	if err := ensureIndexes(ctx, db.Collection("cursos_progresso"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "cliente_uuid", Value: 1}, {Key: "curso_slug", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 12) Chamado
	if err := ensureIndexes(ctx, db.Collection("chamados"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "cliente_uuid", Value: 1}, {Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 13) Recebivel
	if err := ensureIndexes(ctx, db.Collection("recebiveis"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "cliente_uuid", Value: 1}, {Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 14) Pagamento
	if err := ensureIndexes(ctx, db.Collection("pagamentos"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 15) Produto
	if err := ensureIndexes(ctx, db.Collection("produtos"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "familia", Value: 1}, {Key: "slug", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	// 16) Alerta
	if err := ensureIndexes(ctx, db.Collection("alertas"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "cliente_uuid", Value: 1}, {Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "created_at", Value: -1}}},
	}); err != nil {
		return err
	}

	// 17) Notificacao
	if err := ensureIndexes(ctx, db.Collection("notificacoes"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "tipo", Value: 1}, {Key: "created_at", Value: -1}}},
	}); err != nil {
		return err
	}

	// 18) Usuario
	if err := ensureIndexes(ctx, db.Collection("usuarios"), []mongo.IndexModel{
		{Keys: bson.D{{Key: "uuid", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)},
	}); err != nil {
		return err
	}

	return nil
}

func ensureIndexes(ctx context.Context, coll *mongo.Collection, idx []mongo.IndexModel) error {
	if len(idx) == 0 {
		return nil
	}
	_, err := coll.Indexes().CreateMany(ctx, idx)
	return err
}