# 🚀 CRUD de Usuários com Golang + Gin + GORM

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Gonic-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GORM](https://img.shields.io/badge/GORM-ORM-blue?style=for-the-badge)
![SQLite](https://img.shields.io/badge/SQLite-Database-blue?style=for-the-badge&logo=sqlite&logoColor=white)

## ✨ Descrição

Um projeto simples de **CRUD de Usuários** usando **Golang**, **Gin** e **GORM** com **SQLite**, seguindo a arquitetura em camadas:

> **Repository → Service → Controller**

---

## 📦 Tecnologias utilizadas

- ✅ Golang
- ✅ Gin (Framework HTTP)
- ✅ GORM (ORM)
- ✅ SQLite (Banco de dados local)

---

## 📁 Estrutura do Projeto

```bash
users-crud/
├── controller/
│   └── user_controller.go
├── db/
│   └── db.go
├── models/
│   └── user.go
├── repository/
│   └── user_repository.go
├── service/
│   └── user_service.go
├── go.mod
└── main.go
```

---

## ⚙️ Como configurar o projeto

### 1️⃣ Clonar o repositório

```bash
git clone https://github.com/seu-usuario/seu-repo.git
cd users-crud
```

> _(Se estiver usando localmente, ignore essa etapa)_

---

### 2️⃣ Iniciar o módulo Go

```bash
go mod init meu-treino-golang/users-crud
```

---

### 3️⃣ Instalar as dependências

```bash
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/sqlite
```

---

### 4️⃣ Organizar dependências

```bash
go mod tidy
```

---

## ▶️ Como rodar a aplicação

Na raiz do projeto:

```bash
go run main.go
```

Se tudo estiver correto, você verá:

```
Listening and serving HTTP on :8080
```

---

## 🌐 Endpoints disponíveis

### ✅ Criar usuário

```
POST /users
```

**Body (JSON):**

```json
{
  "nome": "João Silva",
  "email": "joao@email.com"
}
```

---

### 📄 Listar usuários

```
GET /users
```

---

### ✏️ Atualizar usuário

```
PUT /users/{id}
```

**Exemplo:**

```
PUT /users/1
```

**Body (JSON):**

```json
{
  "nome": "João Atualizado",
  "email": "joao@email.com"
}
```

---

### 🗑️ Deletar usuário

```
DELETE /users/{id}
```

---

## 🧪 Testando com Postman

### Criar usuário

- Método: `POST`
- URL:

```
http://localhost:8080/users
```

- Body → raw → JSON:

```json
{
  "nome": "Maria Oliveira",
  "email": "maria@email.com"
}
```

---

### Listar usuários

- Método: `GET`
- URL:

```
http://localhost:8080/users
```

---

## 🧪 Testando com `curl`

### Criar usuário

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"nome":"João Silva","email":"joao@email.com"}'
```

---

### 📄 Listar todos os usuários

```bash
curl http://localhost:8080/users
```

---

### ✏️ Atualizar um usuário

```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"nome":"João Atualizado","email":"joao.novo@email.com"}'
```

---

### 🗑️ Deletar um usuário

```bash
curl -X DELETE http://localhost:8080/users/1
```

---

### 🔎 Teste rápido de status

```bash
curl -i http://localhost:8080/users
```

> Se retornar `200 OK`, sua API está funcionando corretamente 🎉

---

## 💾 Banco de dados

O banco utilizado é **SQLite**.

O arquivo é criado automaticamente:

```
users.db
```

Você pode abrir esse arquivo usando:

- **DB Browser for SQLite**
- **SQLiteStudio**
- ou qualquer visualizador de SQLite.

---

## 📚 Arquitetura do Projeto

A aplicação segue o padrão:

```
Controller → Service → Repository → Database
```

- **Controller**: recebe as requisições HTTP (Gin)
- **Service**: aplica regras de negócio
- **Repository**: acessa o banco de dados com GORM

---

## 🛑 Erros comuns

### ❌ 404 Not Found

Verifique a rota:

```
/users
```

ou

```
/api/users
```

---

### ❌ Porta já em uso

```
listen tcp :8080: bind: address already in use
```

Finalize o processo anterior ou altere a porta no `main.go`.

---

## 🧾 Licença

Projeto livre para fins de estudo e aprendizado 📚
