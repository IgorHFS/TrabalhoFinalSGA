package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"sga/handlers"
	"sga/repositories"
	"sga/services"
)

func Register(router *gin.Engine, db *pgxpool.Pool, version string) {
	salas := handlers.SalaHandler{Service: services.SalaService{Repo: repositories.SalaRepository{DB: db}, Alocacoes: repositories.AlocacaoRepository{DB: db}}}
	alunos := handlers.AlunoHandler{Service: services.AlunoService{Repo: repositories.AlunoRepository{DB: db}}}
	turmas := handlers.TurmaHandler{Service: services.TurmaService{Repo: repositories.TurmaRepository{DB: db}}}
	matriculas := handlers.MatriculaHandler{Service: services.MatriculaService{DB: db}}
	alocacoes := handlers.AlocacaoHandler{Service: services.AlocacaoService{DB: db}}
	v1 := router.Group("/api/v1")
	v1.GET("/health", handlers.HealthHandler{Version: version}.Check)
	v1.POST("/salas", salas.Create)
	v1.GET("/salas", salas.List)
	v1.GET("/salas/:id/agenda", salas.Agenda)
	v1.POST("/alunos", alunos.Create)
	v1.GET("/alunos", alunos.List)
	v1.GET("/alunos/:id", alunos.Get)
	v1.POST("/turmas", turmas.Create)
	v1.GET("/turmas", turmas.List)
	v1.POST("/turmas/:id/matriculas", matriculas.Create)
	v1.GET("/turmas/:id/matriculas", matriculas.List)
	v1.POST("/turmas/:id/alocacoes", alocacoes.Create)
}
