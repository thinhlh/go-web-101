# Run with whole stack
run-local:
	@docker compose -f deployment/docker-compose.yml --env-file .env --profile app --profile dependencies up --build

# Run dependencies only
pre-run:
	@docker compose --profile dependencies -f deployment/docker-compose.yml --env-file .env up --build

# Run dependencies only
clear:
	@docker compose -f deployment/docker-compose.yml down --rmi local --volumes

# Install Golang dependencies
install:
	@go mod tidy

# Run in Prod mode
.PHONY: services run-product-service run-order-service
services: run-product-service run-order-service

run-product-service:
	@go run ./cmd/product_service/main.go

run-order-service:
	@go run ./cmd/order_service/main.go

# Run in dev mode
.PHONY: services-dev run-order-service-dev run-order-service-dev
services-dev: run-product-service-dev run-order-service-dev

run-product-service-dev:
	@~/go/bin/air -c .air.product.toml

run-order-service-dev:
	@~/go/bin/air -c .air.order.toml

# Build binary
build:
	@go build -o ./bin/server ./cmd/server/main.go

# Test
test:
	@go test -v ./...