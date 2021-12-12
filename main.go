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
        
        fmt.Fprintf(w,"<img src='https://media-exp1.licdn.com/dms/image/C4D22AQF-oJhwPSSQNQ/feedshare-shrink_1280/0/1638953672281?e=1642032000&amp;v=beta&amp;t=veoZsqkTUV9edSrTlyIQJqX4lzaH44kHBShXwVudQrQ'>")
    })

    http.HandleFunc("/bc4m", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    // Prometheus endpoint
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":8081", nil))

}
