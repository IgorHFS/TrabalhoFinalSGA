package services

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sga/models"
)

type MatriculaService struct{ DB *pgxpool.Pool }

func (s MatriculaService) Create(ctx context.Context, turmaID string, in models.MatriculaInput) error {
	if in.AlunoID == "" {
		return BadRequest("aluno_id é obrigatório")
	}
	tx, err := s.DB.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	var turma, aluno bool
	if err = tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM turmas WHERE id=$1 AND ativa) , EXISTS(SELECT 1 FROM alunos WHERE id=$2)", turmaID, in.AlunoID).Scan(&turma, &aluno); err != nil {
		return err
	}
	if !turma {
		return NotFound("turma não encontrada")
	}
	if !aluno {
		return NotFound("aluno não encontrado")
	}
	var already bool
	if err = tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM matriculas WHERE turma_id=$1 AND aluno_id=$2)", turmaID, in.AlunoID).Scan(&already); err != nil {
		return err
	}
	if already {
		return Conflict("aluno já está matriculado nesta turma")
	}
	var capacidade *int
	var total int
	err = tx.QueryRow(ctx, `SELECT s.capacidade, (SELECT COUNT(*) FROM matriculas WHERE turma_id=$1)::int FROM alocacoes a JOIN salas s ON s.id=a.sala_id WHERE a.turma_id=$1 FOR UPDATE`, turmaID).Scan(&capacidade, &total)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}
	if capacidade != nil && total+1 > *capacidade {
		return Unprocessable("capacidade da sala seria excedida")
	}
	var conflito bool
	err = tx.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM matriculas m JOIN alocacoes a ON a.turma_id=m.turma_id JOIN alocacoes nova ON nova.turma_id=$1 WHERE m.aluno_id=$2 AND a.dia_semana=nova.dia_semana AND nova.hora_inicio<a.hora_fim AND nova.hora_fim>a.hora_inicio)`, turmaID, in.AlunoID).Scan(&conflito)
	if err != nil {
		return err
	}
	if conflito {
		return Conflict("conflito de agenda do aluno")
	}
	if _, err = tx.Exec(ctx, "INSERT INTO matriculas (turma_id,aluno_id) VALUES ($1,$2)", turmaID, in.AlunoID); err != nil {
		return err
	}
	return tx.Commit(ctx)
}
func (s MatriculaService) List(ctx context.Context, turmaID string) ([]models.Aluno, error) {
	var exists bool
	if err := s.DB.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM turmas WHERE id=$1)", turmaID).Scan(&exists); err != nil {
		return nil, err
	}
	if !exists {
		return nil, NotFound("turma não encontrada")
	}
	rows, err := s.DB.Query(ctx, "SELECT a.id,a.nome,a.email FROM alunos a JOIN matriculas m ON m.aluno_id=a.id WHERE m.turma_id=$1 ORDER BY a.nome", turmaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Aluno])
}
