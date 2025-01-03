package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"sort"
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
	parts := strings.Split(strings.TrimPrefix(url, "https://"), "/")
	return strings.TrimPrefix(parts[0], "www.")
}

func (us *URLShortener) handleShorten(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	us.mutex.Lock()
	defer us.mutex.Unlock()

	// Check if URL already exists in cache
	if shortURL, exists := us.urlCache[req.URL]; exists {
		json.NewEncoder(w).Encode(ShortenResponse{ShortURL: shortURL})
		return
	}

	// Generate new short URL
	shortURL := us.generateShortURL(req.URL)
	us.urls[shortURL] = req.URL
	us.urlCache[req.URL] = shortURL

	json.NewEncoder(w).Encode(ShortenResponse{ShortURL: shortURL})
}

func (us *URLShortener) handleRedirect(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimPrefix(r.URL.Path, "/")
	us.mutex.RLock()
	longURL, exists := us.urls[shortURL]
	us.mutex.RUnlock()

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}

func (us *URLShortener) handleMetrics(w http.ResponseWriter, r *http.Request) {
	us.mutex.RLock()
	defer us.mutex.RUnlock()

	var domains []DomainStat
	for domain, count := range us.stats {
		domains = append(domains, DomainStat{Domain: domain, Count: count})
	}

	sort.Slice(domains, func(i, j int) bool {
		return domains[i].Count > domains[j].Count
	})

	if len(domains) > 3 {
		domains = domains[:3]
	}

	json.NewEncoder(w).Encode(MetricsResponse{TopDomains: domains})
}
