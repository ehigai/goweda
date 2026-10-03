package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)


type Condition struct {
	TempC string `json:"temp_C"`
	FeelsLikeC string 
	Humidity string `json:"humidity"`
	WeatherDesc []WeatherDescription `json:"weatherDesc"`
}

type WeatherDescription struct {
	Value string `json:"value"`
}

type WeatherData struct {
	Current []Condition `json:"current_condition"`
}

func main() {
	// Define the city flag with a default value
	city := flag.String("city", "Lagos", "The city to get weather for")
	flag.Parse()

	// Replace spaces with plus sign (+)
	formatedCity := strings.ReplaceAll(*city, " ", "+")

	url := fmt.Sprintf("https://wttr.in/%s?format=j1", formatedCity)

	fmt.Println("Fetching weather details for: ", *city)


	// Hit the url endpoint
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to reach the API: %v\n", err)
		os.Exit(1) // exit the program
	}
	// close connection
	defer response.Body.Close()

	// If we did'nt get a 200 OK response
	if response.StatusCode != http.StatusOK {
		fmt.Printf("API returned an error: %d %s\n", response.StatusCode, response.Status)
		os.Exit(1)
	}

	// Read data stream from response body
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Failed to read the response: %v\n", err)
		os.Exit(1)
	}

	var weather WeatherData

	// Convert the json bytes into a struct
	err = json.Unmarshal(responseBodyBytes, &weather) // we must pass a pointer so it can modify our variable
	if err != nil {
		fmt.Printf("Failed to parse JSON: %v\n", err)
		os.Exit(1)
	}

	if len(weather.Current) > 0 {
		current := weather.Current[0]

		fmt.Printf("\nWeather in %s:\n", *city)
		fmt.Printf("Temperature: %sdeg (Feels like %sdeg)\n", current.TempC, current.FeelsLikeC)
		fmt.Printf("Humidity: %s%%\n", current.Humidity)
		fmt.Printf("Description: %s\n", current.WeatherDesc[0].Value)
		os.Exit(0)
	}

	fmt.Printf("Could not find weather codition for city: %s", *city)
	
}
