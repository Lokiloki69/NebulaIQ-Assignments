package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// High-cardinality dimensions
	methods := []string{"GET", "POST"}
	statuses := []string{"200", "400", "500"}

	// Generate 2000 unique endpoints
	endpoints := make([]string, 2000)
	for i := range endpoints {
		endpoints[i] = fmt.Sprintf("/api/user/%d", i)
	}

	// Counter with 3 labels → VERY HIGH CARDINALITY
	requests := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "High-cardinality HTTP requests generated for testing",
		},
		[]string{"method", "endpoint", "status"},
	)

	// Metrics handler
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Metrics exposed at http://localhost:8000/metrics")

	// Run metrics endpoint
	go func() {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			log.Fatal("Failed to start metrics HTTP server:", err)
		}
	}()

	// Generate ~1000 requests/sec
	log.Println("Generating synthetic metrics...")
	ticker := time.NewTicker(1 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		method := methods[rand.Intn(len(methods))]
		endpoint := endpoints[rand.Intn(len(endpoints))]
		status := statuses[rand.Intn(len(statuses))]

		requests.WithLabelValues(method, endpoint, status).Inc()
	}
}
