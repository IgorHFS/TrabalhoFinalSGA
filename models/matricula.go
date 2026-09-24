package models

type MatriculaInput struct {
	AlunoID string `json:"aluno_id" binding:"required"`
}
