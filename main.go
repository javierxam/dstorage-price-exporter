package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"
	"time"
	"flag"

)

// startMonitor refreshes the prices metrics periodically as defined by refreshRate
func startMonitor(refreshRate time.Duration) {
	for range time.Tick(time.Minute * 1) {
		callClient()
	}
}

func main() {
	port := flag.Int("port", 8004, "Port to serve Prometheus Metrics on")
	flag.Parse()

	// Set the metrics initially before starting the monitor and HTTP server
	// If you don't do this all the metrics start with a "0" until they are set
	callClient()

	// start the metrics collector
	go startMonitor(time.Duration(1))

	// This section will start the HTTP server and expose
	// any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
