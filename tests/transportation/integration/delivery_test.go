package dashboard_test

import (
	"encoding/json"
	"farmerapp/internal/transportation/delivery"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"
)

var BASE_URL string

func loadDotEnv() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		fmt.Errorf("Error loading .env file")
	}
}

func loadEnvVariables() {
	BASE_URL = os.Getenv("BASE_URL")
	if BASE_URL == "" {
		fmt.Errorf("BASE_URL in .env file is missing")
	}
}

func TestGetActiveDeliveries(t *testing.T) {

	// TODO:- Mock data to be first pushed into the db if that's how we are going to test it

	loadDotEnv()
	loadEnvVariables()

	resp, err := http.Get(BASE_URL + "transportation/delivery/active-deliveries")
	assert.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body map[string][]delivery.Delivery
	err = json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)

	// expected := []string{"Active Order A", "Active Order B"}
	// expected := []string(nil)
	// assert.Equal(t, expected, body["deliveries"])

}
