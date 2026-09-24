package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sga/models"
)

type AlunoRepository struct{ DB *pgxpool.Pool }

func (r AlunoRepository) Create(ctx context.Context, a models.Aluno) error {
	_, err := r.DB.Exec(ctx, "INSERT INTO alunos (id,nome,email) VALUES ($1,$2,$3)", a.ID, a.Nome, a.Email)
	return err
}
func (r AlunoRepository) List(ctx context.Context) ([]models.Aluno, error) {
	rows, err := r.DB.Query(ctx, "SELECT id,nome,email FROM alunos ORDER BY nome")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Aluno])
}
func (r AlunoRepository) Get(ctx context.Context, id string) (models.Aluno, error) {
	var aluno models.Aluno
	err := r.DB.QueryRow(ctx, "SELECT id,nome,email FROM alunos WHERE id=$1", id).Scan(&aluno.ID, &aluno.Nome, &aluno.Email)
	return aluno, err
}
