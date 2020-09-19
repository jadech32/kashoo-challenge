CURRENT_DIR = $(shell pwd)
##########################################################################################
# Display usage output
##########################################################################################
all: swagger-docs build swagger-codegen

.PHONY: help
help:
	@echo "Geolocation API Makefile"
	@echo "Builds:"
	@echo
	@echo "  make all                 	Makes swagger docs, builds the binary and client."
	@echo "  make build                 Builds the binary package."
	@echo "  make swagger-docs          Runs echo-swagger on the project for the swagger endpoint."
	@echo "  make swagger-codegen   	Builds the client via swagger-codegen"


.PHONY: build
build:
	go build -o server $(CURRENT_DIR)/cmd/main.go
.PHONY: swagger-docs
swagger-docs:
	swag init -dir "./cmd"

.PHONY: swagger-codegen
swagger-codegen:
	swagger-codegen generate -l go -i ./docs/swagger.json -o ./client

