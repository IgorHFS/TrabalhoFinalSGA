package services

import (
	"context"
	"sga/models"
	"sga/repositories"
)

type TurmaService struct{ Repo repositories.TurmaRepository }

func (s TurmaService) Create(ctx context.Context, in models.Turma) error {
	if in.ID == "" || in.Nome == "" || in.Disciplina == "" || in.Professor == "" {
		return BadRequest("id, nome, disciplina e professor são obrigatórios")
	}
	if err := s.Repo.Create(ctx, in); err != nil {
		if isUnique(err) {
			return Conflict("turma já cadastrada")
		}
		return err
	}
	return nil
}
func (s TurmaService) List(ctx context.Context) ([]models.Turma, error) { return s.Repo.List(ctx) }
