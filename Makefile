SHELL := /bin/bash
BIN := build/pulsa
STATIC_BIN := build/${BUILD_NAME}
IMAGE_NAME ?= pulsa
TAG ?= latest
PORT ?= 8080

current_dir = $(shell pwd)

ENVS := GO_PRIVATE=gitlab.com/msstoci/ CGO_ENABLED=0 GOOS=linux GODEBUG=amd64
FLAGS := -tags=go_json,netgo,nomsgpack -ldflags='-s -w -extldflags "-static"'


$(BIN):
	git config --global url."git@gitlab.com:".insteadOf "https://gitlab.com/"
	mkdir -p build/
	go mod tidy 
	# $(ENVS) go build $(FLAGS) -o $(BIN)

static:
	$(ENVS) go build $(FLAGS) -o $(STATIC_BIN)

lint:
	@echo "Run linter..."
	@golangci-lint run


fmt:
	@echo "Formatting code style..."
	gofumpt -l -w \
		config/.. \
		internal/.. \
		pkg/..
	@echo "[DONE] Formatting code style..."

imports:
	@echo "Formatting imports..."
	goimports -w -local bitbucket.org/efishery/efishery-fish-service-user,bitbucket.org/efishery/go-efishery \
		config/.. \
		internal/.. \
		pkg/..
	@echo "[DONE] Formatting imports..."

format:
	@$(MAKE) fmt
	@$(MAKE) imports