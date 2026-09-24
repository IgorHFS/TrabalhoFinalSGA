package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sga/models"
)

type AlocacaoRepository struct{ DB *pgxpool.Pool }

func (r AlocacaoRepository) AgendaSala(ctx context.Context, salaID string) ([]models.Alocacao, error) {
	rows, err := r.DB.Query(ctx, "SELECT turma_id,sala_id,dia_semana,to_char(hora_inicio, 'HH24:MI') hora_inicio,to_char(hora_fim, 'HH24:MI') hora_fim FROM alocacoes WHERE sala_id=$1 ORDER BY dia_semana,hora_inicio", salaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return pgx.CollectRows(rows, pgx.RowToStructByName[models.Alocacao])
}
