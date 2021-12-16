package main

import (
	"net/http"

	"github.com/nelkinda/health-go"
	"github.com/nelkinda/health-go/checks/sysinfo"
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
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}
