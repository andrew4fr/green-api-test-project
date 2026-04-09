.PHONY: generate install-codegen lint lint-fix

LOCAL_BIN=$(CURDIR)/bin
MODELS_DIR := models
HANDLERS_DIR := handlers
SWAGGER_FILE := api/swagger.yaml

OAPI_CODEGEN := $(shell go env GOPATH)/bin/oapi-codegen

GOLANGCI_BIN=$(LOCAL_BIN)/golangci-lint
$(GOLANGCI_BIN):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCAL_BIN) v2.5.0

install-codegen:
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

generate: install-codegen
	$(OAPI_CODEGEN) -generate types -package models \
		$(SWAGGER_FILE) > $(MODELS_DIR)/models.go

	$(OAPI_CODEGEN) -generate chi-server,types  -package handlers \
		$(SWAGGER_FILE) > $(HANDLERS_DIR)/handlers.go

lint: $(GOLANGCI_BIN)
	$(GOENV) $(GOLANGCI_BIN) run ./... -v

lint-fix: $(GOLANGCI_BIN)
	$(GOENV) $(GOLANGCI_BIN) run ./... --fix -v
