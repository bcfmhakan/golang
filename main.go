package main

import (
	"net/http"

	"github.com/nelkinda/health-go"
	"github.com/nelkinda/health-go/checks/sysinfo"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	h := health.New(
		health.Health{
			Version:   "1",
			ReleaseID: "1.0.0",
		},
		sysinfo.Health(),
	)
	http.HandleFunc("/health", h.Handler)
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}
