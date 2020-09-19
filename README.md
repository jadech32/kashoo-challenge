## Geolocation API

A basic geolocation RESTful API with two endpoints, using echo and Swagger.

## Installation & Usage

### Installation

To install all required packages: `go mod download`.

Make sure you have swagger installed.

NOTE: The `Makefile` assumes you have `swagger-codegen` installed via homebrew. Feel free to edit the Makefile target if your codegen is called something else.

### API Key

A ipgeolocation.io API key is required. Create a `.env` file in the root directory that contains the `API_KEY` token. Currently an `.env` file is required for the program to run (no other way to load env variables)

### How To Run

Since a Swagger/OpenAPI specification is required as one of the API's endpoints, run `make swagger-docs` to create RESTful API documentation for our API.

1. Build the binary with `make build` or `make all` to also create the client.
2. Run binary with `./server`


### Endpoints

**GET /geolocate/<ip_address>**

IP geolocation that queries an IP address and returns location data

**GET /swagger/index.html**

Displays Swagger/OpenAPI specification of the endpoint above