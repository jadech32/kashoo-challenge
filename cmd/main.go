package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	_ "github.com/jadech32/kashoo-challenge/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Geolocation API
// @version 1.0
// @description Queries and returns location data for IP addresses

// @host localhost:8080
// @basePath /

// @schemes http https
// @produce	json

func main() {
	// Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// See if API_KEY exists.
	_, ok := os.LookupEnv("API_KEY")
	if !ok {
		log.Fatal("API_KEY environment variable does not exist.")
	}

	e := echo.New()

	// Routes
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/geolocate/:ip", geolocateHandler)

	// Start
	e.Logger.Fatal(e.Start(":8080"))

}

// geolocateHandler godoc
// @Summary Get geolocation of an IP
// @Description Queries ipgeolocation.io for a valid IP address' geolocation.
// @ID geolocate
// @Produce json
// @Param ip path string true "IP Address"
// @Success 200 {object} map[string]interface{} "Successful Operation"
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /geolocate/{ip} [get]
func geolocateHandler(c echo.Context) error {
	ip := c.Param("ip")

	// Request IP Geolocation Data
	u, err := url.Parse("https://api.ipgeolocation.io/ipgeo")
	if err != nil {
		log.Fatal(err)
	}

	name := os.Getenv("API_KEY")

	queryString := u.Query()
	queryString.Add("ip", ip)
	queryString.Add("apiKey", name)

	u.RawQuery = queryString.Encode()

	res, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	// ipgeolocation 403: If the queried IP address or domain name is not valid.
	if res.StatusCode == http.StatusForbidden {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Invalid IP Address",
		})
	}

	// ipgeolocation 4xx: Errors with ipgeolocation
	if res.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to Fetch",
		})
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal to map[string]interface{}
	var result map[string]interface{}
	json.Unmarshal(bodyBytes, &result)

	return c.JSON(http.StatusOK, result)
}

// HTTPError struct
type HTTPError struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"Invalid IP Address"`
}
