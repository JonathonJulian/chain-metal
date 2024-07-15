package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"custom-geth-exporter/metrics"
	"custom-geth-exporter/ui"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ipcPath := os.Getenv("GETH_IPC_PATH")
	httpURL := os.Getenv("GETH_HTTP_URL")
	listenerPort := os.Getenv("METRICS_LISTENER_PORT")
	httpFallback := os.Getenv("HTTP_FALLBACK") == "true"

	err := metrics.Init(ipcPath, httpURL, httpFallback)
	if err != nil {
		log.Fatalf("Failed to initialize Ethereum client: %v", err)
	}

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			if err := metrics.UpdatePeerMetrics(); err != nil {
				log.Printf("Failed to update peer metrics: %v", err)
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/ui", ui.ServeUI)
	http.HandleFunc("/ws", ui.ServeRPCData)

	log.Fatal(http.ListenAndServe(":"+listenerPort, nil))
}
