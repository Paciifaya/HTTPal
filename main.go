package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	method := flag.String("method", "GET", "HTTP method to use")
	url := flag.String("url", "", "URL to request")
	flag.Parse()

	if *url == "" {
		fmt.Println("Usage: fetch -method GET -url https://example.com")
		os.Exit(1)
	}

	request, err := http.NewRequest(
		strings.ToUpper(*method),
		*url,
		nil,
	)

	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}

	request.Header.Set("User-Agent", "fetch-cli/1.0")

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}

	defer response.Body.Close()

	fmt.Printf("HTTP status: %s\n", response.Status)

	for name, values := range response.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}

	fmt.Println()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("could not read response body: %v", err)
	}

	fmt.Println(string(body))
}
