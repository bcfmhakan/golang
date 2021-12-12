package main

import 
    "html/template"
    "fmt"
    "html"
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

    
    // An HTML template

        
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        
        fmt.Fprintf(w,"#BC4M", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/bc4m", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    // Prometheus endpoint
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":8081", nil))

}
