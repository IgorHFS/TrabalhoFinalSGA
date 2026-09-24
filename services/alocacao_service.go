package services

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sga/models"
	"time"
)

type AlocacaoService struct{ DB *pgxpool.Pool }

func (s AlocacaoService) Create(ctx context.Context, turmaID string, in models.AlocacaoInput) error {
	ini, fim, err := parseHorario(in.HoraInicio, in.HoraFim)
	if err != nil {
		return BadRequest("horários devem estar no formato HH:MM e início deve ser anterior ao fim")
	}
	_ = ini
	_ = fim
	tx, err := s.DB.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.Serializable})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	var turma, sala bool
	if err = tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM turmas WHERE id=$1 AND ativa), EXISTS(SELECT 1 FROM salas WHERE id=$2)", turmaID, in.SalaID).Scan(&turma, &sala); err != nil {
		return err
	}
	if !turma {
		return NotFound("turma não encontrada")
	}
	if !sala {
		return NotFound("sala não encontrada")
	}
	var alunos, capacidade int
	if err = tx.QueryRow(ctx, "SELECT (SELECT COUNT(*) FROM matriculas WHERE turma_id=$1)::int,(SELECT capacidade FROM salas WHERE id=$2)", turmaID, in.SalaID).Scan(&alunos, &capacidade); err != nil {
		return err
	}
	if alunos > capacidade {
		return Unprocessable("capacidade da sala é insuficiente")
	}
	var conflSala bool
	if err = tx.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM alocacoes WHERE turma_id<>$1 AND sala_id=$2 AND dia_semana=$3 AND $4::time<hora_fim AND $5::time>hora_inicio)", turmaID, in.SalaID, in.DiaSemana, in.HoraInicio, in.HoraFim).Scan(&conflSala); err != nil {
		return err
	}
	if conflSala {
		return Conflict("conflito de agenda da sala")
	}
	var conflAluno bool
	if err = tx.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM matriculas nova JOIN matriculas antiga ON antiga.aluno_id=nova.aluno_id JOIN alocacoes a ON a.turma_id=antiga.turma_id WHERE nova.turma_id=$1 AND antiga.turma_id<>$1 AND a.dia_semana=$2 AND $3::time<a.hora_fim AND $4::time>a.hora_inicio)`, turmaID, in.DiaSemana, in.HoraInicio, in.HoraFim).Scan(&conflAluno); err != nil {
		return err
	}
	if conflAluno {
		return Conflict("há aluno com conflito de agenda")
	}
	_, err = tx.Exec(ctx, "INSERT INTO alocacoes (turma_id,sala_id,dia_semana,hora_inicio,hora_fim) VALUES ($1,$2,$3,$4,$5) ON CONFLICT (turma_id) DO UPDATE SET sala_id=EXCLUDED.sala_id,dia_semana=EXCLUDED.dia_semana,hora_inicio=EXCLUDED.hora_inicio,hora_fim=EXCLUDED.hora_fim", turmaID, in.SalaID, in.DiaSemana, in.HoraInicio, in.HoraFim)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}
func parseHorario(a, b string) (time.Time, time.Time, error) {
	x, e := time.Parse("15:04", a)
	if e != nil {
		return x, time.Time{}, e
	}
	y, e := time.Parse("15:04", b)
	if e != nil || !x.Before(y) {
		return x, y, errors.New("horário inválido")
	}
	return x, y, nil
}
