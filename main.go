package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "github.com/etherlabsio/healthcheck/v2"
    "github.com/etherlabsio/healthcheck/v2/checkers"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hi")
    })

    // Prometheus endpoint
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":8081", nil))
    r := mux.NewRouter()
    r.Handle("/healthcheck", healthcheck.Handler(

        // WithTimeout allows you to set a max overall timeout.
        healthcheck.WithTimeout(5*time.Second),

        // Checkers fail the status in case of any error.
        healthcheck.WithChecker(
            "heartbeat", checkers.Heartbeat("$PROJECT_PATH/heartbeat"),
        ),

        healthcheck.WithChecker(
            "database", healthcheck.CheckerFunc(
                func(ctx context.Context) error {
                    return db.PingContext(ctx)
                },
            ),
        ),

        // Observers do not fail the status in case of error.
        healthcheck.WithObserver(
            "diskspace", checkers.DiskSpace("/var/log", 90),
        ),
    ))
}
