package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        
        fmt.Fprintf(w, "<img src='https://media-exp1.licdn.com/dms/image/C4D22AQF-oJhwPSSQNQ/feedshare-shrink_1280/0/1638953672281?e=1642032000&v=beta&t=veoZsqkTUV9edSrTlyIQJqX4lzaH44kHBShXwVudQrQ' >", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    // Prometheus endpoint
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":8081", nil))

}
