# SGA — Sistema de Gestão de Alocação

API REST em Go para cadastro de salas, alunos e turmas, com matrícula e alocação de horários.

## Executar

Preencha as variáveis `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` e `DB_SSLMODE` no arquivo `.env` (ou no ambiente) e execute:

```powershell
go run ./api
```

As tabelas e o índice necessário são criados automaticamente no PostgreSQL. A API usa `SERVER_PORT` (padrão `8080`) e `APP_VERSION` (padrão `1.0.0`).

## Endpoints

| Método | Rota | Finalidade |
| --- | --- | --- |
| GET | `/api/v1/health` | Saúde, horário UTC e versão |
| POST / GET | `/api/v1/salas` | Criar e listar salas |
| GET | `/api/v1/salas/:id/agenda` | Consultar agenda da sala |
| POST / GET | `/api/v1/alunos` | Criar e listar alunos |
| GET | `/api/v1/alunos/:id` | Consultar aluno |
| POST / GET | `/api/v1/turmas` | Criar e listar turmas |
| POST / GET | `/api/v1/turmas/:id/matriculas` | Matricular e listar alunos da turma |
| POST | `/api/v1/turmas/:id/alocacoes` | Alocar ou realocar uma turma |

Exemplo de alocação:

```json
{
  "sala_id": "LAB-01",
  "dia_semana": 1,
  "hora_inicio": "08:00",
  "hora_fim": "10:00"
}
```

`dia_semana` usa 1 para segunda-feira e 7 para domingo. Requisições inválidas retornam `400`; recursos ausentes, `404`; duplicidade ou conflito de agenda, `409`; e capacidade insuficiente, `422`.
