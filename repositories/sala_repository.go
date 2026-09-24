package repositories

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sga/models"
)

type SalaRepository struct{ DB *pgxpool.Pool }

func (r SalaRepository) Create(ctx context.Context, s models.Sala) error {
	_, err := r.DB.Exec(ctx, "INSERT INTO salas (id,nome,capacidade,recursos) VALUES ($1,$2,$3,$4)", s.ID, s.Nome, s.Capacidade, s.Recursos)
	return err
}
func (r SalaRepository) List(ctx context.Context) ([]models.Sala, error) {
	rows, err := r.DB.Query(ctx, "SELECT id,nome,capacidade,recursos FROM salas ORDER BY nome")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	salas, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Sala])
	return salas, err
}
