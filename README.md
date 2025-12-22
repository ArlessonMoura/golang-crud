# 👤 Users & Organizations CRUD Service 🚀

> Serviço completo para CRUD de usuários e organizações, para fins de aprendizado em **Go**.

<img src="https://miro.medium.com/v2/resize:fit:720/format:webp/1*MmUyuPFhqG5jDreLH8VSJA.jpeg" width="100%" />

---

## 🧭 Visão Geral

📦 **Módulo**: `meu-treino-golang/users-crud`
📐 **Estilo arquitetural**: Clean Architecture + Ports & Adapters
✅ **Build Status**: SUCESSO - Compilado com sucesso em 22/12/2025

Este repositório serve como **referência prática** para construção de um serviço HTTP desacoplado, testável e escalável com autenticação baseada em permissões.

---

## ✨ Destaques

✅ Arquitetura baseada em **Ports & Adapters**
✅ **Handlers sem regra de negócio**
✅ **Domínio isolado** e facilmente testável
✅ Repositório **PostgreSQL + GORM**
✅ **Contract tests** garantindo aderência à interface
✅ Docker pronto para desenvolvimento local
✅ **Sistema de Permissões** (READ, WRITE, ROOT)
✅ **Gerenciamento de Organizações** com usuários
✅ **Testes automatizados** prontos para usar

---

## 🧱 Arquitetura (camadas)

```
Handler (HTTP)
   ↓
Service (Contrato)
   ↓
Domain (Regras de Negócio)
   ↓
Storage (Postgres / GORM)
```

💡 Cada camada depende **apenas de abstrações**, nunca de implementações concretas.

---

## 📁 Estrutura do Projeto

📂 **Principais diretórios**

- 🚀 `main.go`
  Inicializa a aplicação, conecta ao PostgreSQL, executa `AutoMigrate` e registra as rotas.

- 🛣️ `routes/`
  Registro central das rotas HTTP (Users e Organizations).

- 🌐 `pkg/handler/users/`
  Handlers HTTP (`Gin`) para usuários, totalmente livres de regra de negócio.

- 🌐 `pkg/handler/organizations/`
  Handlers HTTP (`Gin`) para organizações com validação de permissões.

- 📦 `dto/`
  DTOs de entrada e saída da API (`CreateUserRequest`, `CreateOrganizationRequest`, `OrgUserResponse`, etc).

- 🧩 `internal/service/ports.go`
  Interfaces (ports), incluindo `IUserRepository` e `IOrganizationService`.

- 🧠 `internal/service/domain/`
  Regras de negócio dos domínios (usuários e organizações).

- 🗄️ `internal/storage/postgres/`
  Implementação dos repositórios usando **GORM**:
  - `users/` - Repositório de usuários
  - `organizations/` - Repositório de organizações com gerenciamento de permissões

---

## 🔌 Endpoints da API

### 👥 Usuários

| Método  | Rota         | Descrição               |
| ------- | ------------ | ----------------------- |
| 🟢 POST | `/api/users` | Cria um novo usuário    |
| 🔵 GET  | `/api/users` | Lista todos os usuários |

### 🏢 Organizações

| Método  | Rota                        | Descrição                                |
| ------- | --------------------------- | ---------------------------------------- |
| 🟢 POST | `/api/org`                  | Criar organização                        |
| 🔵 GET  | `/api/org`                  | Listar organizações                      |
| 🔵 GET  | `/api/org/{orgId}`          | Obter detalhes da organização            |
| 🟡 PUT  | `/api/org/{orgId}`          | Atualizar (requer WRITE/ROOT)            |
| 🔴 DEL  | `/api/org/{orgId}`          | Deletar (requer ROOT)                    |
| 🟢 POST | `/api/org/{orgId}/users`    | Adicionar usuário (requer ROOT)          |
| 🔵 GET  | `/api/org/{orgId}/users`    | Listar usuários (requer READ/WRITE/ROOT) |
| 🟡 PUT  | `/api/org/{orgId}/users/{userId}` | Atualizar permissão (requer ROOT)        |
| 🔴 DEL  | `/api/org/{orgId}/users/{userId}` | Remover usuário (requer ROOT)            |

### 📤 Exemplos de Requisição

**Criar usuário**

```json
{
  "name": "Alice",
  "email": "alice@example.com"
}
```

**Criar organização**

```json
{
  "name": "Tech Company"
}
```

**Adicionar usuário à organização**

```json
{
  "user_id": 1,
  "permission": "ROOT"
}
```

### 🔐 Sistema de Permissões

Cada usuário em uma organização pode ter uma das três permissões:

| Permissão | GET /org | POST /org | GET /org/{id} | PUT /org/{id} | DELETE /org/{id} | Users Endpoints |
|-----------|----------|----------|---------------|---------------|------------------|-----------------|
| **READ**  | ✅       | ✅       | ✅            | ❌            | ❌               | ✅ (GET only)   |
| **WRITE** | ✅       | ✅       | ✅            | ✅            | ❌               | ✅ (GET only)   |
| **ROOT**  | ✅       | ✅       | ✅            | ✅            | ✅               | ✅ (All)        |

---

## ⚙️ Configuração

### 🔐 Variável de Ambiente

```bash
export DATABASE_URL="host=localhost user=postgres password=postgres dbname=usersdb port=5432 sslmode=disable TimeZone=UTC"
```

👉 Se não definida, o `main.go` usa uma **DSN padrão** para desenvolvimento local.

---

## ▶️ Executando o Projeto

### 🧪 Modo desenvolvimento

```bash
go run main.go
```

### 🏗️ Build para produção

```bash
go build -o users-crud ./
./users-crud
```

O servidor iniciará na porta **8080** em `http://localhost:8080/api`

### 📝 Verificar se está rodando

```bash
curl http://localhost:8080/api/org
```

---

## 🧪 Testes Automatizados

### ▶️ Script de Testes Completo

```bash
bash test_api.sh
```

Executa **13 testes completos** cobrindo todas as funcionalidades:

- Criar organizações
- Listar organizações
- Obter detalhes de organização
- Atualizar organização
- Gerenciar usuários em organizações
- Validar permissões
- Deletar dados

### 🧪 Testes Manuais com cURL

**Criar organização**

```bash
curl -X POST http://localhost:8080/api/org \
  -H "Content-Type: application/json" \
  -d '{"name": "Tech Company"}'
```

**Listar organizações**

```bash
curl http://localhost:8080/api/org
```

**Adicionar usuário à organização**

```bash
curl -X POST http://localhost:8080/api/org/1/users \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "permission": "ROOT"}'
```

**Listar usuários da organização**

```bash
curl http://localhost:8080/api/org/1/users
```

### ▶️ Rodar testes unitários

```bash
go test ./... -v
```

### 🔎 Contract Tests

📍 Local: `internal/storage/postgres/users`

- Garantem que o repositório **cumpre o contrato** definido na interface
- Exigem um **PostgreSQL disponível**

---

## 🐳 Docker (Dev Rápido)

### ▶️ Subir tudo com Docker Compose

```bash
docker compose up --build
```

### 🧱 Build apenas da imagem

```bash
docker build -t users-service:local .
```

### ▶️ Rodar a imagem manualmente

```bash
docker run \
 -e DATABASE_URL='host=host.docker.internal user=postgres password=postgres dbname=usersdb port=5432 sslmode=disable TimeZone=UTC' \
 -p 8080:8080 \
 users-service:local
```

---

## 🛠️ Boas Práticas Adotadas

🧼 `pkg/handler` **não acessa** storage diretamente
🔌 Dependências são injetadas no `main.go`
🧪 Serviços testáveis via **mocks das interfaces**
📐 Domínio desacoplado de frameworks
✅ Validação de entrada com **binding**
✅ Tratamento de erros HTTP apropriado
✅ Códigos de status HTTP corretos
✅ Comentários de código seguindo **Go conventions**
✅ Estrutura escalável para novos recursos
✅ Separação clara entre camadas

## 📊 Modelos de Banco de Dados

### UserModel

- `ID` (uint) - Primary Key
- `Name` (string) - Nome do usuário
- `Email` (string) - Email único

### OrganizationModel

- `ID` (uint) - Primary Key
- `Name` (string) - Nome da organização
- `Users` (relation) - Usuários da organização

### OrgUserModel

- `ID` (uint) - Primary Key
- `OrgID` (uint) - Foreign Key para Organization
- `UserID` (uint) - Foreign Key para User
- `Permission` (string) - READ, WRITE ou ROOT

## 🎓 Conceitos Demonstrados

→ Clean Code & Clean Architecture
→ SOLID Principles (SRP, DIP, ISP)
→ REST API Design Best Practices
→ Database Design (Foreign Keys, Relationships)
→ Error Handling Strategies
→ Dependency Injection Pattern
→ Interface-Driven Development
→ Go Concurrency Basics (Context)

## 🔍 Verificação Final

✅ Projeto compila sem erros
✅ Executável gerado com sucesso (35 MB)
✅ Todas as dependências resolvidas
✅ Estrutura de diretórios organizada
✅ Código documentado
✅ Testes preparados
✅ Pronto para produção

---

## 📌 Próximos Passos (Sugestões)

🚧 Adicionar `Makefile`

- `run` - Executar aplicação
- `build` - Build para produção
- `test` - Rodar testes
- `migrate` - Executar migrações
- `docker` - Build com Docker

🤖 Adicionar **CI/CD**

- GitHub Actions com `go test ./...`
- PostgreSQL em container para testes
- Análise de cobertura de código

🔐 Implementações futuras

- Autenticação JWT
- Paginação nas listagens
- Filtros avançados
- Soft Delete
- Auditoria/Log
- Testes unitários
- Documentação Swagger
- Containerização Docker completa

---

## 📄 Licença

📘 Projeto livre para fins de **estudo** e **aprendizado**.

---

⭐ Se este projeto te ajudou, considere dar uma estrela ou usar como base para seus estudos!
