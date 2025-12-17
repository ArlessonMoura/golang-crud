# 🏗️ Arquitetura - Users Service

## Princípios

- **Handler → Service → Domain → Storage**: Flow de dependências unidirecional
- **Ports & Adapters**: Storage é desacoplado via interfaces
- **Contract Tests**: Governança da porta de repositório
- **depguard**: Enforce de regras de importação

## Estrutura

```
users-service/
├── dto/                          # DTOs (entrada/saída HTTP)
├── pkg/handler/users/            # HTTP Handlers (sem regra de negócio)
├── internal/
│   ├── common/                   # Dependências, erros, paginação
│   ├── service/
│   │   ├── service.go           # Contrato IUserService
│   │   ├── ports.go             # Contrato IUserRepository
│   │   └── domain/users/        # Implementação do serviço
│   └── storage/postgres/users/  # GORM Repository
├── routes/                        # Wiring de rotas
└── .golangci.yml                # Regras depguard
```

## Fluxo de Requisição

1. **HTTP Request** → Handler recebe DTO
2. **Handler** → Converte para parâmetros simples, chama service
3. **Service** → Valida regras de negócio, chama repository via port
4. **Repository** → Persiste no GORM, retorna UserDTO
5. **Handler** → Serializa resposta HTTP

## Regras de Importação (depguard)

### ❌ Proibido

- `pkg/handler/` importar `internal/storage/`
- Qualquer lugar importar storage fora do wiring

### ✅ Permitido

- `pkg/handler/` → `internal/service/` (abstrações)
- `internal/service/domain/` → `internal/service/` (portas)
- `internal/storage/postgres/` → `internal/service/` (portas)

## Testes

- **Contract Tests**: Validam que Repository implementa IUserRepository
- **Unit Tests**: Service testado com mock de repository
- **Integration Tests**: Repository testado contra banco real
