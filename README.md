# Users CRUD Service

Repositório de referência para um serviço HTTP minimalista de CRUD de usuários, implementado com arquitetura limpa (Handler → Service → Domain → Storage).

**Módulo**: `meu-treino-golang/users-crud` (veja [go.mod](go.mod)).

Visão geral (resumo):

- `pkg/handler` — camada HTTP (rotas e handlers).
- `internal/service` — contratos (interfaces) e DTOs.
- `internal/service/domain/users` — regras de negócio e validações.
- `internal/storage/postgres/users` — adapter GORM que implementa a interface de repositório.
- `internal/common` — provisionamento de dependências (ex.: DB).
- `routes` — função para registrar rotas no servidor HTTP.

Estrutura relevante (trechos importantes):

- `main.go` — inicia a aplicação, conecta ao PostgreSQL, executa `AutoMigrate` e registra rotas ([main.go](main.go)).
- `routes/routes.go` — registra handlers com o `gin.Engine` ([routes/routes.go](routes/routes.go)).
- `pkg/handler/users/` — `InitHandler`, `Handler` e rotas para `/api/users` ([pkg/handler/users/handler.go](pkg/handler/users/handler.go)).
- `dto/` — DTOs HTTP: `CreateUserRequest`, `UserResponse` ([dto/user_dto.go](dto/user_dto.go)).
- `internal/service/ports.go` — contrato do repositório (`IUserRepository`) ([internal/service/ports.go](internal/service/ports.go)).
- `internal/service/service.go` — contrato do serviço (`IUserService`) e `UserDTO` ([internal/service/service.go](internal/service/service.go)).
- `internal/service/domain/users/service.go` — implementação do serviço de domínio ([internal/service/domain/users/service.go](internal/service/domain/users/service.go)).
- `internal/storage/postgres/users/repository.go` — implementação GORM do repositório e `UserModel` ([internal/storage/postgres/users/repository.go](internal/storage/postgres/users/repository.go)).

Endpoints principais

- POST `/api/users` — cria usuário. Recebe `CreateUserRequest` (JSON) e retorna `{ "id": <id> }`.
- GET `/api/users` — lista usuários. Retorna array de `UserResponse`.

Configuração e execução

1. Configure a variável de ambiente `DATABASE_URL`. Se não setada, `main.go` usa uma DSN padrão para desenvolvimento local:

```bash
export DATABASE_URL="host=localhost user=postgres password=postgres dbname=usersdb port=5432 sslmode=disable TimeZone=UTC"
```

2. Rodar em modo desenvolvimento:

```bash
go run main.go
```

3. Build para produção:

```bash
go build -o users-service ./
./users-service
```

Testes

- Executar todos os testes (unitários e de pacote):

```bash
go test ./... -v
```

- Contract tests: há testes que validam se o repositório implementa corretamente a porta (`internal/storage/postgres/users/contract_test.go`). Esses testes podem exigir um PostgreSQL disponível quando realizam operações de integração.

Observações de desenvolvimento

- Mantenha a regra de importação: `pkg/handler` nunca deve importar `internal/storage` diretamente. O ponto único de wiring é `main.go`/`routes`/`internal/common`.
- Para testar a camada de serviço, use mocks da interface `internal/service.IUserRepository`.
- As migrations são feitas via `gorm.AutoMigrate` no `main.go` para `UserModel`.

Exemplos rápidos (cURL)

Criar usuário:

```bash
curl -X POST http://localhost:8080/api/users \
 -H "Content-Type: application/json" \
 -d '{"name":"Alice","email":"alice@example.com"}'
```

Listar usuários:

```bash
curl http://localhost:8080/api/users
```

Próximos passos sugeridos

- Adicionar um `Makefile` com alvos: `run`, `build`, `test`, `migrate`.
- Adicionar CI que rode `go test ./...` e execute contract tests contra um Postgres em container.
- Posso adicionar o `Makefile` e o workflow de CI de exemplo se desejar.

---

Licença: Projeto livre para fins de estudo e aprendizado.
