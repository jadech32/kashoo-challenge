## Geolocation API

A basic geolocation RESTful API with two endpoints, using echo and Swagger.

## Installation & Usage

### Installation

To install all required packages: `go mod download`.

Make sure you have swagger installed.

NOTE: The `Makefile` assumes you have `swagger-codegen` installed via homebrew. Feel free to edit the Makefile target if your codegen is called something else.

### API Key

A ipgeolocation.io API key is required. Create a `.env` file in the root directory that contains the `API_KEY` token, or otherwise (e.g. Kubernetes ConfigMap, etc.)

### How To Run

#### As an executable

1. Build the binary with `make build`
2. Run binary with `./server`
