package main

import (
	"fmt"
	"net/http"
)

func main() {
	shortener := NewURLShortener()

	http.HandleFunc("/shorten", shortener.handleShorten)
	http.HandleFunc("/metrics", shortener.handleMetrics)
	http.HandleFunc("/", shortener.handleRedirect)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
