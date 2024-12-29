package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"sync"
)

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

type MetricsResponse struct {
	TopDomains []DomainStat `json:"top_domains"`
}

type DomainStat struct {
	Domain string `json:"domain"`
	Count  int    `json:"count"`
}

func (us *URLShortener) generateShortURL(url string) string {
	hash := sha256.Sum256([]byte(url))
	return base64.URLEncoding.EncodeToString(hash[:8])
}

func (us *URLShortener) extractDomain(url string) string {
	parts := strings.Split(strings.TrimPrefix(url, "http://"), "/")
	return strings.TrimPrefix(parts[0], "www.")
}

func (us *URLShortener) handleShorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	us.mutex.Lock()
	defer us.mutex.Unlock()

	shortURL := us.generateShortURL(req.URL)
	us.urls[shortURL] = req.URL
	us.urlCache[req.URL] = shortURL

	//updating my domain statistics
	domain := us.extractDomain(req.URL)
	us.stats[domain]++
	json.NewEncoder(w).Encode(ShortenResponse{ShortURL: shortURL})
}
