package main

import "sync"

type URLShortener struct {
	urls     map[string]string
	stats    map[string]int
	urlCache map[string]string
	mutex    sync.RWMutex
}

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		urls:     make(map[string]string),
		stats:    make(map[string]int),
		urlCache: make(map[string]string),
	}
}
