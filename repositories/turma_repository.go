package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sga/models"
)

type TurmaRepository struct{ DB *pgxpool.Pool }

func (r TurmaRepository) Create(ctx context.Context, t models.Turma) error {
	_, err := r.DB.Exec(ctx, "INSERT INTO turmas (id,nome,disciplina,professor) VALUES ($1,$2,$3,$4)", t.ID, t.Nome, t.Disciplina, t.Professor)
	return err
}
func (r TurmaRepository) List(ctx context.Context) ([]models.Turma, error) {
	rows, err := r.DB.Query(ctx, `SELECT t.id,t.nome,t.disciplina,t.professor,COUNT(m.aluno_id)::int quantidade_alunos,EXISTS(SELECT 1 FROM alocacoes a WHERE a.turma_id=t.id) alocada FROM turmas t LEFT JOIN matriculas m ON m.turma_id=t.id GROUP BY t.id ORDER BY t.nome`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Turma])
}
