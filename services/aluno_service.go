package services

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"sga/models"
	"sga/repositories"
)

type AlunoService struct{ Repo repositories.AlunoRepository }

func (s AlunoService) Create(ctx context.Context, in models.Aluno) error {
	if in.ID == "" || in.Nome == "" || in.Email == "" {
		return BadRequest("id, nome e email são obrigatórios")
	}
	if err := s.Repo.Create(ctx, in); err != nil {
		if isUnique(err) {
			return Conflict("aluno ou email já cadastrado")
		}
		return err
	}
	return nil
}
func (s AlunoService) List(ctx context.Context) ([]models.Aluno, error) { return s.Repo.List(ctx) }
func (s AlunoService) Get(ctx context.Context, id string) (models.Aluno, error) {
	a, err := s.Repo.Get(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return a, NotFound("aluno não encontrado")
	}
	return a, err
}
