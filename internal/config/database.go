package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	_ = godotenv.Load()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), value("DB_PORT", "5432"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), value("DB_SSLMODE", "disable"))
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("criar pool PostgreSQL: %w", err)
	}
	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("conectar ao PostgreSQL: %w", err)
	}
	if err = migrate(ctx, pool); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}

func value(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func migrate(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS salas (id TEXT PRIMARY KEY, nome TEXT NOT NULL, capacidade INTEGER NOT NULL CHECK (capacidade > 0), recursos TEXT[] NOT NULL DEFAULT '{}');
		CREATE TABLE IF NOT EXISTS alunos (id TEXT PRIMARY KEY, nome TEXT NOT NULL, email TEXT NOT NULL UNIQUE);
		CREATE TABLE IF NOT EXISTS turmas (id TEXT PRIMARY KEY, nome TEXT NOT NULL, disciplina TEXT NOT NULL, professor TEXT NOT NULL, ativa BOOLEAN NOT NULL DEFAULT TRUE);
		CREATE TABLE IF NOT EXISTS matriculas (turma_id TEXT NOT NULL REFERENCES turmas(id) ON DELETE CASCADE, aluno_id TEXT NOT NULL REFERENCES alunos(id) ON DELETE CASCADE, PRIMARY KEY (turma_id, aluno_id));
		CREATE TABLE IF NOT EXISTS alocacoes (turma_id TEXT PRIMARY KEY REFERENCES turmas(id) ON DELETE CASCADE, sala_id TEXT NOT NULL REFERENCES salas(id), dia_semana SMALLINT NOT NULL CHECK (dia_semana BETWEEN 1 AND 7), hora_inicio TIME NOT NULL, hora_fim TIME NOT NULL, CHECK (hora_inicio < hora_fim));
		CREATE INDEX IF NOT EXISTS idx_alocacoes_sala_dia ON alocacoes (sala_id, dia_semana);`)
	if err != nil {
		return fmt.Errorf("migrar banco de dados: %w", err)
	}
	return nil
}
