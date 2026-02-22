include .env
export

PG_DSN=postgres://$(PG_USER):$(PG_PASS)@$(PG_HOST):$(PG_PORT)/$(PG_DB_NAME)?sslmode=disable

compose-up:
	docker compose -f deployments/docker-compose.yml --env-file .env up

compose-down:
	docker compose -f deployments/docker-compose.yml down

compose-build:
	docker compose -f deployments/docker-compose.yml --env-file .env build

generate-sqlc:
	sqlc generate

migrate-pgsql-create:
	$(eval NAME ?= todo)
	goose -dir deployments/migrations postgres $(PG_DSN) create init sql

generate-proto:
	protoc \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		proto/**/*.proto

clean-proto:
	@echo "Cleaning generated proto files..."
	rm -f proto/transactions/*.pb.go
	rm -f proto/clients/*.pb.go
	@echo "✓ Cleaned"

install-proto-tools:
	@echo "Installing protoc plugins..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "✓ Installed"

tidy:
	(cd pkg && go mod tidy)
	(cd proto && go mod tidy)
	(cd services/api-gateway && go mod tidy)
	(cd services/notification-service && go mod tidy)