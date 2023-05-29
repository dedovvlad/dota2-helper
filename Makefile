GO               = go
VERSION          := $(shell git describe --tags)
DATE             := $(shell date +%FT%T%z)
GIT_COMMIT       := $(shell git rev-parse HEAD)
GIT_BRANCH       := $(shell git rev-parse --abbrev-ref HEAD)
M                = $(shell printf "\033[34;1m>>\033[0m")
GOBIN			 ?= $(PWD)/bin
MIGRATIONS_DIR	 = ./db/migrations/
SQLC_TAG		 =1.17.0
SQLC_PLATFORM 	 := $(if $(SQLC_PLATFORM),$(SQLC_PLATFORM),"darwin_amd64")

.PHONY: all
all: build test

.PHONY: build
build: build-service

.PHONY: build-service
build-service: ## Build binary
	$(info $(M) building service...)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(LDFLAGS) -o $(GOBIN)/service ./cmd/*.go

watch: install-tools ; ## Run binaries that rebuild themselves on changes
	$(info $(M) run...)
	@$(GOBIN)/air -c .air.conf

.PHONY: fmt
fmt: ## Format code
	$(info $(M) running gofmt...)
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		$(GO) fmt $$d/*.go || ret=$$? ; \
		done ; exit $$ret

.PHONY: install-tools
install-tools: install-goose install-sqlc install-lint
	@GOBIN=$(GOBIN) $(GO) install -mod=readonly github.com/golang/mock/mockgen@latest

install-goose:
	@GOBIN=$(GOBIN) $(GO) install -mod=readonly github.com/pressly/goose/v3/cmd/goose@v3.7.0

install-lint:
	@GOBIN=$(GOBIN) $(GO) install -mod=readonly github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2

install-sqlc:
	@mkdir -p bin
	@echo "Downloading sqlx ${SQLC_TAG} / ${SQLC_PLATFORM}"
	@wget -c "https://github.com/kyleconroy/sqlc/releases/download/v${SQLC_TAG}/sqlc_${SQLC_TAG}_${SQLC_PLATFORM}.tar.gz" -O - | tar -xz  -C ${GOBIN}
	@chmod +x ${GOBIN}/sqlc

.PHONY: install-tools-gimps
install-tools-gimps:
ifeq ($(wildcard $(GOBIN)/gimps),)
	@echo "Downloading gimps"
	@GOBIN=$(GOBIN) $(GO) install -mod=readonly go.xrstf.de/gimps@latest
endif

.PHONY: imports-fix
imports-fix: install-tools-gimps ; ## Fix imports
	$(info $(M) fixing imports...)
	@$(GOBIN)/gimps -config .gimps.yaml .

.PHONY: lint
lint: install-lint ## Run linters
	$(info $(M) running linters...)
	@$(GOBIN)/golangci-lint run --timeout 5m0s ./...

.PHONY: lintfix
lintfix: install-lint ## Try to fix linter issues
	$(info $(M) fixing linter issues...)
	@$(GOBIN)/golangci-lint run --fix --verbose --timeout 2m0s ./... 2>&1 | \
		awk 'BEGIN{FS="="} /Fix/ { print $$3}' | \
		awk 'BEGIN{FS=","} {print " * ", $$1, $$2, $$8, $$9, $$10, $$11}' | \
		sed 's/\\"/"/g' | sed -e 's/&result.Issue{//g' | sed 's/token.Position//'

.PHONY: test
test: ## Run all tests
	$(info $(M) running tests...)
	@$(GO) test ./... -v -cover

.PHONY: db-migrate
db-migrate: ## Run migrate command
	$(info $(M) running DB migrations...)
	@$(GOBIN)/goose -dir "$(MIGRATIONS_DIR)" postgres "$(STORAGE_MIGRATION_DSN)" $(filter-out $@,$(MAKECMDGOALS))

.PHONY: db-create-migration
db-create-migration: ## Create a new database migration file
	$(info $(M) creating DB migration...)
	@$(GOBIN)/goose -dir "$(MIGRATIONS_DIR)" create $(filter-out $@,$(MAKECMDGOALS)) sql

.PHONY: db-apply-migrations
db-apply-migrations: ## Run database migrations
	$(info $(M) applying DB migrations...)
	@$(GOBIN)/goose -v -dir "$(MIGRATIONS_DIR)" postgres "$(STORAGE_MIGRATION_DSN)" up
	@$(GOBIN)/goose -v postgres "$(STORAGE_MIGRATION_DSN)" status

.PHONY: generate
generate: ## Run go generate
	$(info $(M) generating...)
	@$(GO) generate ./...

.PHONY: vet
vet: ## Run go vet
	$(info $(M) vetting source...)
	@go vet ./...

.PHONY: clean
clean:
	rm -rf $(GOBIN)

.PHONY: $(GOBIN)
$(GOBIN):
	@mkdir -p $(GOBIN)

.PHONY: start
start: ## Start dev environment
	$(info $(M) starting dev environment...)
	@docker-compose down
	@cp ~/.netrc .netrc
	@docker-compose up $(filter-out $@,$(MAKECMDGOALS))

.PHONY: stop
stop: ; $(info $(M) stopping development Docker env...)
	@echo Stopping containers...
	@docker-compose down
	@echo OK

.PHONY: test-createdb
test-createdb:
	@docker-compose exec db createdb --username=postgres --owner=postgres test

.PHONY: test-dropdb
test-dropdb:
	@docker-compose exec db dropdb --username=postgres test

help:                   ##Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

%:
	@:
