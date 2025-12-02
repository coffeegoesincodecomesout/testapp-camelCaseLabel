package main

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "log"
    "math/rand"
    "net/http"
    "time"
)

var (
    counter = promauto.NewCounter(prometheus.CounterOpts{
      Name: "ping_summary_duration_ms",
      Help: "Counting the total number of requests handled",
    })
    
    // must be registered
    gauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
      Name: "ping_gauge",
      Help: "testing camelCase labels",
    }, []string{"node", "nameSpace"})
)

func recordMetrics() {
    go func() {
      for {
        counter.Inc()
        gauge.WithLabelValues("node-1", "/namespace-b").Set(rand.Float64())
        time.Sleep(time.Second * 5)
      }
    }()
}

func init() {
    prometheus.MustRegister(gauge)
}

func main() {
    recordMetrics()
    
    srv := http.NewServeMux()
    srv.Handle("/metrics", promhttp.Handler())
    
    if err := http.ListenAndServe(":8090", srv); err != nil {
        log.Fatalf("unable to start server: %v", err)
    }
}
