package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "github.com/prometheus/client_golang/tree/master/prometheus/promhttp"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

   	// Prometheus endpoint
	router.Path("/prometheus").Handler(promhttp.Handler())

    log.Fatal(http.ListenAndServe(":8081", nil))

}
