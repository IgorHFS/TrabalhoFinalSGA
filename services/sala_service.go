package services

import (
	"context"
	"sga/models"
	"sga/repositories"
)

type SalaService struct {
	Repo      repositories.SalaRepository
	Alocacoes repositories.AlocacaoRepository
}

func (s SalaService) Create(ctx context.Context, in models.Sala) error {
	if in.ID == "" || in.Nome == "" || in.Capacidade < 1 {
		return BadRequest("id, nome e capacidade positiva são obrigatórios")
	}
	if err := s.Repo.Create(ctx, in); err != nil {
		if isUnique(err) {
			return Conflict("sala já cadastrada")
		}
		return err
	}
	return nil
}
func (s SalaService) List(ctx context.Context) ([]models.Sala, error) { return s.Repo.List(ctx) }
func (s SalaService) Agenda(ctx context.Context, id string) ([]models.Alocacao, error) {
	var exists bool
	err := s.Repo.DB.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM salas WHERE id=$1)", id).Scan(&exists)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, NotFound("sala não encontrada")
	}
	return s.Alocacoes.AgendaSala(ctx, id)
}
