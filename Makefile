# Run with whole stack
run-local:
	@docker compose -f deployment/docker-compose.yml --env-file env/.env.dev up --build -d

# Run dependencies only
pre-run:
	@docker compose --profile dependencies -f deployment/docker-compose.yml --env-file env/.env.dev up --build

# Install Golang dependencies
install:
	@go mod tidy

run:
	@go run ./cmd/server/main.go

# Run in dev mode
run-dev:
	@~/go/bin/air -c .air.toml

# Build binary
build:
	@go build -o ./bin/server ./cmd/server/main.go

# Test
test:
	@go test -v ./...