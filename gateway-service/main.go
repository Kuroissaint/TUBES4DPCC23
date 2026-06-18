package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	var targetURL string

	if strings.HasPrefix(r.URL.Path, "/location") {
		targetURL = "http://location-service-service:8002"
		// Kita TIDAK LAGI memotong /location, biarkan utuh!
	} else if strings.HasPrefix(r.URL.Path, "/dispatch") {
		targetURL = "http://dispatch-service-service:8003"
		// Kita TIDAK LAGI memotong /dispatch, biarkan utuh sampai ke belakang!
	} else {
		http.Error(w, "Service Not Found", http.StatusNotFound)
		return
	}

	target, _ := url.Parse(targetURL)
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", proxyHandler)
	log.Println("API Gateway running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}