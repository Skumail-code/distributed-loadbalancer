package main

import (
    "log"
    "net/http"
    "time"
)

func main() {
    servers := []string{"server1:8080", "server2:8080"}
    lb := NewLoadBalancer(servers)
    go lb.PeriodicHealthCheck(30 * time.Second)

    metrics := &Metrics{}
    http.Handle("/metrics", metrics)
    http.Handle("/", lb)

    log.Println("Load balancer starting...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
