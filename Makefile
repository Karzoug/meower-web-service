# ==============================================================================
# Define dependencies

SERVICE_VERSION            := 0.1.0
SERVICE_NAME               := web-service
OUTBOX_VERSION             := 0.1.0
BUILD_VERSION              ?= $(shell git symbolic-ref HEAD 2> /dev/null | cut -b 12-)_$(shell git log --pretty=format:%h -1)
BUILD_DATE                 ?= $(shell date +%FT%T%z)

BASE_IMAGE                 := meower
IMAGE_SERVICE              := $(BASE_IMAGE)/service/web:$(SERVICE_VERSION)

OAPI_CODEGEN_VERSION       := 2.4.1
GOLANGCI_LINT_VERSION      := 1.61.0
BUF_VERSION                := 1.46.0
PROTOC_GEN_GO_VERSION 	   := 1.35.1
PROTOC_GEN_GO_GRPC_VERSION := 1.5.1

MAIN_PACKAGE_PATH          := ./cmd/
BINARY_NAME                := web_service
TEMP_DIR                   := /var/tmp/meower/web
TEMP_BIN                   := ${TEMP_DIR}/bin
PROJECT_PKG                := github.com/Karzoug/meower-web-service

LDFLAGS += -s -w
LDFLAGS += -X ${PROJECT_PKG}/pkg/buildinfo.buildVersion=${BUILD_VERSION} -X ${PROJECT_PKG}/pkg/buildinfo.buildDate=${BUILD_DATE} -X ${PROJECT_PKG}/pkg/buildinfo.serviceVersion=$(SERVICE_VERSION)

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	@test -z "$(shell git status --porcelain)"


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: run quality control checks
.PHONY: audit
audit: test fmt lint
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)" 
	go vet ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

## fmt: format .go files
.PHONY: fmt
fmt:
	go run golang.org/x/tools/cmd/goimports@latest -local=${PROJECT_PKG} -l -w  .
	go run mvdan.cc/gofumpt@latest -l -w  .

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

## lint: run linters
.PHONY: lint
lint:
	$(TEMP_BIN)/golangci-lint run ./...

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## tidy: tidy modfiles and format .go files
.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

## build: build the application
.PHONY: build
build:
	go build -ldflags "${LDFLAGS}" -o ${TEMP_BIN}/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## generate: generate all necessary code
.PHONY: generate
generate:
	$(TEMP_BIN)/oapi-codegen --config=oapi_server.config.yaml api/openapi/v1/api.yaml
	$(TEMP_BIN)/oapi-codegen --config=oapi_models.config.yaml api/openapi/v1/api.yaml
	$(TEMP_BIN)/buf generate

## clean: clean all temporary files
.PHONY: clean
clean:
	rm -rf $(TEMP_DIR)

# ==============================================================================
# Install dependencies

## dev-install-deps: install dependencies with fixed versions in a temporary directory
dev-install-deps:
	GOBIN=$(TEMP_BIN) go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v${OAPI_CODEGEN_VERSION}
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(TEMP_BIN) v${GOLANGCI_LINT_VERSION}
	GOBIN=$(TEMP_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GEN_GO_VERSION}
	GOBIN=$(TEMP_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOC_GEN_GO_GRPC_VERSION}
	GOBIN=$(TEMP_BIN) go install github.com/bufbuild/buf/cmd/buf@v$(BUF_VERSION)

# ==============================================================================
# Building containers

## service: build the service image
.PHONY: service
service:
	docker build \
		-f build/dockerfile.service \
		-t $(IMAGE_SERVICE) \
		--build-arg BUILD_REF=$(BUILD_VERSION) \
		--build-arg BUILD_DATE=$(BUILD_DATE) \
		--build-arg VERSION=$(SERVICE_VERSION) \
		--build-arg PROJECT_PKG=$(PROJECT_PKG) \
		--build-arg SERVICE_NAME=$(SERVICE_NAME) \
		.

# ------------------------------------------------------------------------------

dev-compose-up:
	IMAGE_SERVICE=$(IMAGE_SERVICE) docker compose -f deploy/docker/dev-compose.yaml  up -d

dev-compose-down:
	IMAGE_SERVICE=$(IMAGE_SERVICE) docker compose -f deploy/docker/dev-compose.yaml down
