package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/centrifugal/gocent/v3"
	"log"
	"net/http"
	"time"
)

func main() {
	// Handle the root URL ("/") by serving the index.html file
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	go sender()

	go gocentSender()

	// Start the web server on port 8080
	log.Println("Starting web server on http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Data represents the JSON payload
type Data struct {
	Channel string         `json:"channel"`
	Data    map[string]int `json:"data"`
}

func gocentSender() {

	c := gocent.New(gocent.Config{
		Addr: "http://localhost:8000/api",
		Key:  "my_api_key",
	})

	i := 1000000
	for {
		i++
		data, err := json.Marshal(map[string]int{"value": i})
		if err != nil {
			log.Fatalf("Error marshalling data: %v", err)
		}
		result, err := c.Publish(context.Background(), "channel", data)
		if err != nil {
			log.Fatalf("Error calling publish: %v", err)
		}
		log.Printf("Publish into channel %s successful, stream position {offset: %d, epoch: %s}", "channel", result.Offset, result.Epoch)

		// Sleep for a specified duration before sending the next request
		time.Sleep(2 * time.Second) // Send a request every 5 seconds
	}

}

// sendPostRequest sends a POST request to the specified URL with the given data
func sendPostRequest(url string, apiKey string, payload Data) error {
	// Convert payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with status code: %d", resp.StatusCode)
	} else {
		log.Println("Request successful")
	}

	return nil
}

func sender() {
	url := "http://localhost:8000/api/publish"
	apiKey := "my_api_key"

	// Send the request in a loop
	var i int
	for {
		i++
		payload := Data{
			Channel: "channel",
			Data: map[string]int{
				"value": i,
			},
		}
		err := sendPostRequest(url, apiKey, payload)
		if err != nil {
			log.Printf("Error sending request: %v", err)
		}

		// Sleep for a specified duration before sending the next request
		time.Sleep(5 * time.Second) // Send a request every 5 seconds
	}
}
