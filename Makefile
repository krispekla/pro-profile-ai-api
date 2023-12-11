
db_con ?= "postgres://$$PPAI_API_DB_USER:$$PPAI_API_DB_PASSWORD@$$PPAI_API_DB_HOST:$$PPAI_API_DB_PORT/$$PPAI_API_DB_NAME?sslmode=disable"

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go test -race -buildvcs -vet=off ./...


# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## build: build the cmd/api application
.PHONY: build
build:
	go build -o=./bin/api ./cmd/api

## run: run the cmd/api application
.PHONY: run
run: build
	./bin/api --env="./.env" \

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "./bin/api" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

## migrate/create: create new migration
## Usage: make migrate/create name="your_migration_name_here"
migrate/create:
	migrate create -ext sql -dir ./migrations -seq $(name)

## migrate/up: run up migration
.PHONY: migrate/up
migrate/up:
	migrate -path ./migrations -database "$(db_con)" up $(n)

## migrate/down: run down migration
.PHONY: migrate/down
migrate/down:
	migrate -path ./migrations -database "$(db_con)" down $(n)
    # Command to rollback migrations

## migrate/print: print migrate version
.PHONY: migrate/print
migrate/print:
	migrate -path ./migrations -database "$(db_con)" version

## migrate/go: run migration to version
.PHONY: migrate/go
migrate/go:
	migrate -path ./migrations -database "$(db_con)" goto $(v)