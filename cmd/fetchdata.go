package cmd

import (
	"io"
	"log"
	"net/http"
)

func FetchActivityData(url string) []byte {

	// Make the request
	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		log.Println("request failed")
	}

	// Get the response
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println("not a valid response")
	}

	// Convert to []byte
	responseToBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("unable to convert to []bytes")
	}

	return responseToBytes

}
