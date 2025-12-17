<img src="https://miro.medium.com/v2/resize:fit:720/format:webp/1*MmUyuPFhqG5jDreLH8VSJA.jpeg" width="100%" />

# 👤 Users CRUD Service 🚀

> Serviço para CRUD de usuários, para fins de aprendizado em **Go**.
>
> **Fluxo arquitetural:** `Handler → Service → Domain → Storage`

---

## 🧭 Visão Geral

📦 **Módulo**: `meu-treino-golang/users-crud`
📐 **Estilo arquitetural**: Clean Architecture + Ports & Adapters

Este repositório serve como **referência prática** para construção de um serviço HTTP simples, desacoplado e testável.

---

## ✨ Destaques

✅ Arquitetura baseada em **Ports & Adapters**
✅ **Handlers sem regra de negócio**
✅ **Domínio isolado** e facilmente testável
✅ Repositório **PostgreSQL + GORM**
✅ **Contract tests** garantindo aderência à interface
✅ Docker pronto para desenvolvimento local

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
  Registro central das rotas HTTP.

- 🌐 `pkg/handler/users/`
  Handlers HTTP (`Gin`), totalmente livres de regra de negócio.

- 📦 `dto/`
  DTOs de entrada e saída da API (`CreateUserRequest`, `UserResponse`).

- 🧩 `internal/service/ports.go`
  Interfaces (ports), incluindo `IUserRepository`.

- 🧠 `internal/service/domain/users/`
  Regras de negócio do domínio de usuários.

- 🗄️ `internal/storage/postgres/users/`
  Implementação do repositório usando **GORM** + testes de contrato.

---

## 🔌 Endpoints da API

| Método  | Rota         | Descrição               |
| ------- | ------------ | ----------------------- |
| 🟢 POST | `/api/users` | Cria um novo usuário    |
| 🔵 GET  | `/api/users` | Lista todos os usuários |

### 📤 Criar usuário

```json
{
  "name": "Alice",
  "email": "alice@example.com"
}
```

📥 **Resposta**

```json
{ "id": 1 }
```

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
go build -o users-service ./
./users-service
```

---

## 🧪 Testes

### ▶️ Rodar todos os testes

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

---

## 📌 Próximos Passos (Sugestões)

🚧 Adicionar `Makefile`

- `run`
- `build`
- `test`
- `migrate`

🤖 Adicionar **CI**

- `go test ./...`
- Postgres em container para contract tests

👉 Posso adicionar **Makefile** e **workflow de CI** se desejar 😉

---

## 📄 Licença

📘 Projeto livre para fins de **estudo** e **aprendizado**.

---

⭐ Se este projeto te ajudou, considere usar como base ou dar uma estrela no repositório!
