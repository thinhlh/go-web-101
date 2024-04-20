# Golang microservice Template

## Prerequisite

- Docker/Docker compose
- Golang
- Makefile (optional)
- Copy env
  ```
  cp .env.example .env
  ```

## Boot

1. Local (Boot whole stack)
   ```
   make run
   ```
2. Dev
   ```
   make pre-run
   make -j2 services-dev
   ```
3. Prod
   ```
   make pre-run
   make -j2 services
   ```

> A production ready GoLabg template with Domain Driven Design architecture and predefined configurations

# Features

- [Environment configuration](#environment-configuration)
- [Built in exception filter](#built-in-exception-filter)
- [Base RestAPI response](#base-universial-restapi-response-with-openapi-support)
- [SQL database connection](#sql-database-connection-with-typeorm)
- [Commithook & commitlint](#commithook--commitlint)
- [Error handling](#error-handling)
- [Database migration](#database-migration)
- [Body compression](#body-compression)

## Environment configuration

## Built-in exception filter

## Base universial RestAPI response with OpenAPI support

## SQL Database connection with TypeORM

## Commithook & commitlint

## Error handling

## Database migration

# TODO items

- Server Configs
  - [x] [Cors](./internal/core/middlewares/middlewares.go)
  - [x] [Healthcheck](./internal/core/middlewares/middlewares.go) (Heartbeat middleware)
  - [x] [Profiller](./internal/core/middlewares/middlewares.go)
  - [x] [Api Rate limiter](./internal/core/middlewares/middlewares.go)
- Application Configs
  - [x] Domain Driven architecture
  - [x] [Application load config from file & env](./internal/core/config/config.go)
  - [x] Microservices project architecture
  - [x] Base response for RestAPI
- Deployments
  - [x] [Dockerfile & Docker compose](./deployment/)
- Others

  - [x] [Dev Mode - Reload on changes](./.air.product.toml)

  - [x] Database migration

    - Using gorm AutoMigration

    ```diff
    - Does not have versioning strategy
    - Does not support rollback
    - Does not isolate between infrastructure & application
    + Suggest using other tools that run migration before running application
    // https://www.reddit.com/r/golang/comments/12mypec/how_do_you_handle_migrations/
    ```

    - Current solution: Using golang-migration. Need to add a docker-compose version

    ```diff
    - No tracking & schema -> sql generation
    - Have to write SQL up & down manually
    + Separation of concern
    ```

    Usage

    ```bash
    make migration-create migration_msg="create_table_product" # Create new migration then filling up & down sql script
    make migration-up step= # Up migration. step=N/"" for N step up or all up
    make migration-down step= # Down migration. step=N/"-all" for N step down or all down
    ```

  - [ ] Linting & Commit Lint
  - [ ] Business Error handling
  - [ ] Sample CRUD
  - [ ] CI/CD pipeline
  - [ ] CQRS
  - [ ] BaseModule, BaseCRUD
  - [ ] Pagination
  - [ ] Locking mechanism
  - [ ] Feature flag
  - [ ] HttpServices
  - [ ] Auth
  - [ ] Logging, Tracing & Monitoring
  - [ ] Deployment
  - [ ] Caching
  - [ ] MessageQueue
  - [ ] GraphQL
  - [ ] gRPC
  - [ ] Monorepo
  - [ ] CMS Integration
  - [ ] High-level application architecture docs
