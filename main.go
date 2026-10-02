package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Define the city flag with a default value
	city := flag.String("city", "Lagos", "The city to get weather for")
	flag.Parse()

	// Replace spaces with plus sign (+)
	formatedCity := strings.ReplaceAll(*city, " ", "+")

	url := fmt.Sprintf("https://wttr.in/%s?format=j1", formatedCity)

	fmt.Println("Fetching weather details for: ", *city)
	fmt.Println("API URL: ", url)


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

	// conver bytes to string
	rawJSON := string(responseBodyBytes)

	fmt.Println("\n--- Raw API Response (First 300 characters) ---")
	if len(rawJSON) > 300 {
		fmt.Println(rawJSON[:300] + "...\n")
	} else {
		fmt.Println(rawJSON)
	}
}
