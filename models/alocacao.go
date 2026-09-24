package models

type Alocacao struct {
	TurmaID    string `json:"turma_id" db:"turma_id"`
	SalaID     string `json:"sala_id" db:"sala_id"`
	DiaSemana  int    `json:"dia_semana" db:"dia_semana"`
	HoraInicio string `json:"hora_inicio" db:"hora_inicio"`
	HoraFim    string `json:"hora_fim" db:"hora_fim"`
}

type AlocacaoInput struct {
	SalaID     string `json:"sala_id" binding:"required"`
	DiaSemana  int    `json:"dia_semana" binding:"required,min=1,max=7"`
	HoraInicio string `json:"hora_inicio" binding:"required"`
	HoraFim    string `json:"hora_fim" binding:"required"`
}
