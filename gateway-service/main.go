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

	// 1. Routing untuk Location & Dispatch
	if strings.HasPrefix(r.URL.Path, "/location") {
		targetURL = "http://location-service-service:8002"
	} else if strings.HasPrefix(r.URL.Path, "/dispatch") {
		targetURL = "http://dispatch-service-service:8003"
<<<<<<< HEAD

	// 2. Routing untuk Shop Order & Translog
	} else if strings.HasPrefix(r.URL.Path, "/api/order") {
		targetURL = "http://shop-order-service:8084"
	} else if strings.HasPrefix(r.URL.Path, "/api/translog") {
		targetURL = "http://translog-service:8085"

	// 3. Routing untuk Account Service (mencakup Auth, Driver, dan Customer)
	} else if strings.HasPrefix(r.URL.Path, "/api/auth") || strings.HasPrefix(r.URL.Path, "/api/driver") || strings.HasPrefix(r.URL.Path, "/api/customer") {
		targetURL = "http://account-service:8081"

	// 4. Routing untuk Merchant Service
	} else if strings.HasPrefix(r.URL.Path, "/api/merchant") {
		targetURL = "http://merchant-service:8089"

	// 5. Routing untuk Finance / Wallet Service
	} else if strings.HasPrefix(r.URL.Path, "/api/wallet") {
		targetURL = "http://finance-service:8086"

	// 6. Routing untuk Pricing Service
	} else if strings.HasPrefix(r.URL.Path, "/api/pricing") {
		targetURL = "http://pricing-service:8087"

	// 7. Routing untuk Communication / Chat Service
	} else if strings.HasPrefix(r.URL.Path, "/chat") {
		targetURL = "http://communication-service:8009"

	// 8. Routing untuk Review & Rating Service
	} else if strings.HasPrefix(r.URL.Path, "/reviews") {
		targetURL = "http://reviewrating-service:8008"

	// Jika rute tidak ditemukan di daftar atas
=======
>>>>>>> 6e1565d031e0aa4899b96495fc91c46423db8ade
	} else {
		http.Error(w, "Gateway Error: Service Not Found for this path", http.StatusNotFound)
		return
	}

	// Proses Proxy Request
	target, err := url.Parse(targetURL)
	if err != nil {
		http.Error(w, "Gateway Error: Bad Target URL", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", proxyHandler)
	log.Println("API Gateway running on port 8080 and ready to route all microservices...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}