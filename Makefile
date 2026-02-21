include .env
export

protoDir := "./proto"

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
	rm -f proto/events/*.pb.go
	@echo "✓ Cleaned"

install-proto-tools:
	@echo "Installing protoc plugins..."
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@echo "✓ Installed"