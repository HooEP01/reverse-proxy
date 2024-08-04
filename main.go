package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	targetURL, err := url.Parse("http://example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Create a reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Start the server
	log.Println("Starting reverse proxy server on port 8080...")
	err = http.ListenAndServe(":8080", proxy)
	if err != nil {
		log.Fatal(err)
	}
}
