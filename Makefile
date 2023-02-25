.PHONY: gen

GO_BIN := $(shell echo ${GOPATH}/bin)
OUTPUT := bin/main
OPENAPI_SPEC_DIR := openapi
build:
	@go build -o ${OUTPUT} ./cmd/...
tidy:
	@go mod tidy
install:
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
gen:
	@$(GO_BIN)/oapi-codegen --config=$(OPENAPI_SPEC_DIR)/server.cfg.yml $(OPENAPI_SPEC_DIR)/swagger.yml
	@$(GO_BIN)/oapi-codegen --config=$(OPENAPI_SPEC_DIR)/types.cfg.yml $(OPENAPI_SPEC_DIR)/swagger.yml
version:
	@go version -m ${OUTPUT}
