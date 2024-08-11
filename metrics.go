package main

import (
    "fmt"
    "net/http"
    "sync"
)

type Metrics struct {
    mu                sync.Mutex
    requestsHandled   int
    serversHealthy    int
    responseTimes     []float64
}

func (m *Metrics) IncrementRequests() {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.requestsHandled++
}

func (m *Metrics) UpdateResponseTime(time float64) {
    m.mu.Lock()
    defer m.mu.Unlock()
    m.responseTimes = append(m.responseTimes, time)
}

func (m *Metrics) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    m.mu.Lock()
    defer m.mu.Unlock()

    fmt.Fprintf(w, "Requests Handled: %d\n", m.requestsHandled)
    fmt.Fprintf(w, "Servers Healthy: %d\n", m.serversHealthy)
    fmt.Fprintf(w, "Average Response Time: %f\n", average(m.responseTimes))
}

func average(times []float64) float64 {
    var sum float64
    for _, t := range times {
        sum += t
    }
    return sum / float64(len(times))
}
