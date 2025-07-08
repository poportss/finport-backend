# Caminhos
MAIN_APP=cmd/main.go
MIGRATION_BUILDER=tools/build_migrations.go
I18N_GENERATOR=tools/gen_i18n_maps.go

# Gera a pasta migrations_build/
migrations:
	go run $(MIGRATION_BUILDER)

# Gera arquivos de mensagens traduzidas (i18n)
i18n-build:
	go run $(I18N_GENERATOR)

# Inicia o app
run:
	go run $(MAIN_APP)

# Atalho para desenvolvimento (builda tudo + inicia)
dev: migrations i18n-build run

# Roda os testes
test:
	go test ./...

# Limpa arquivos gerados
clean:
	rm -rf migrations_build

# Roda migrations no banco (necessário o binário do golang-migrate instalado)
migrate-up:
	migrate -path migrations_build/ -database "$$DB_DSN" up

# Rola para trás uma versão
migrate-down:
	migrate -path migrations_build/ -database "$$DB_DSN" down 1
