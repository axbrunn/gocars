include .env

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## The help rule is only on linux
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/api: run the cmd/api application
.PHONY: run/api
run/web:
	@go run ./cmd/web
    # @go run ./cmd/api -db-dsn=${DB_DSN}

.PHONY: db/psql
db/psql:
	psql ${DB_DSN}

.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations -database ${DB_DSN} up


.PHONY: db/migrations/down
db/migrations/down: confirm
	@echo 'Running down migrations...'
	@read -p "How many steps to rollback? " steps; \
	migrate -path ./migrations -database ${DB_DSN} down $$steps
