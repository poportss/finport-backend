# Finport Backend

Backend do sistema de controle financeiro e de investimentos pessoais.

Estruturado em **Clean Architecture** com suporte a:

* Autenticação via JWT
* Multi-idioma (Português, Inglês, Espanhol)
* Migrations por módulo
* API REST (framework: Gin)
* Banco de dados: PostgreSQL (via GORM)

---

## 💪 Tecnologias

* Go 1.21+
* GORM + PostgreSQL
* Gin
* Makefile
* JSON i18n builder

---

## 🚀 Como rodar localmente

1. **Configure o `.env`**

Crie um arquivo `.env` com:

```env
AUTH_SECRET=super-secret-key
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=finport
```

> Use `.env.example` como referência.

---

2. **Rode o banco local com Docker (se desejar)**

```bash
docker run --name finport-db -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=finport -d postgres
```

---

3. **Gere traduções multilíngues**

```bash
make i18n-build
```

---

4. **Rode as migrations e o app**

```bash
make dev
```

---

## 🔪 Testes

```bash
make test
```

---

## 📁 Estrutura de Pastas

```
internal/
├── application/       # Casos de uso e lógica de negócio
├── domain/            # Interfaces e entidades puras
├── infrastructure/    # Repositórios e adaptadores (GORM, etc)
├── interfaces/        # HTTP Handlers, Middlewares
├── services/          # JWT, logging, helpers
└── i18n/              # Traduções geradas automaticamente
```

---

## 🧹 Internacionalização (i18n)

As mensagens estão definidas em `translations/messages.json`.
Use `make i18n-build` para gerar os arquivos `.go` com traduções para:

* Português (`PtMessages`)
* Inglês (`EnMessages`)
* Espanhol (`EsMessages`)

Uso:

```go
msg := i18n.GetMessage(i18n.CreateUserAlias, c)
```

---

## 🏃‍♂️ Versionamento

Versão inicial:

```
v0.1.0 - Estrutura base com JWT, Clean Architecture e i18n
```

---

## 🔐 Autenticação

Rota protegida:

```go
authGroup.Use(middleware.AuthMiddleware(jwtService))
authGroup.GET("/me", MeHandler(userService))
```

---

## 🪰 Comandos Makefile

```bash
make           # builda e roda o backend com migrations
make run       # roda o backend
make migrations  # gera migrations build
make i18n-build  # gera arquivos de tradução
make clean     # limpa pastas geradas
```

---

## 📄 Licença

MIT
