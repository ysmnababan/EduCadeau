package helper

import (
	"donation_service/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func GetPriceFromAddress(from, to string) (float64, error) {
	origin := from
	destination := to

	if origin == "" || destination == "" {
		return 0, fmt.Errorf("origin or destination cannot be blank")
	}

	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")

	// URL encode the origin and destination parameters
	origin = url.QueryEscape(origin)
	destination = url.QueryEscape(destination)

	// Construct the URL for the Google Maps Distance Matrix API
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?origins=%s&destinations=%s&key=%s", origin, destination, apiKey)

	// Make a GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed getting url")
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Error response from API: %s", string(body))
		return 0, fmt.Errorf("status internal server error")
	}

	// Decode the JSON response
	var distanceResponse models.DistanceResponse
	err = json.NewDecoder(resp.Body).Decode(&distanceResponse)
	if err != nil {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Failed to parse distance response: %s", string(body))
		return 0, fmt.Errorf("failed to parse distance response")

	}

	// Check if the API responded with an error status
	if distanceResponse.Status != "OK" {
		return 0, fmt.Errorf("distance calculation failed")
	}

	// Extract the distance text or value as needed
	// var distanceText string //6.4 km
	var distanceValue int
	var deliveryCost int

	if len(distanceResponse.Rows) > 0 && len(distanceResponse.Rows[0].Elements) > 0 {

		distanceValue = distanceResponse.Rows[0].Elements[0].Distance.Value
	}

	if distanceValue == 0 {
		return 0, fmt.Errorf("distance cannot be 0")
	}
	deliveryCost = CalculateShippingCost(distanceValue)

	return float64(deliveryCost), nil
}
