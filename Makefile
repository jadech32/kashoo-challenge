CURRENT_DIR = $(shell pwd)
##########################################################################################
# Display usage output
##########################################################################################
all: image docker-push

.PHONY: help
help:
	@echo "Geolocation API Makefile"
	@echo "Builds:"
	@echo
	@echo "  make build                 Builds the binary package."
	@echo "  make swagger	         	Builds swagger documentation."
	
.PHONY: build
build:
	go build -o server $(CURRENT_DIR)/cmd/main.go

.PHONY: swagger
swagger:
	swagger-codegen generate -l go -i ./docs/swagger.json -o ./client

