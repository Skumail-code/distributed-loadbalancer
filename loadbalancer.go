package main

import (
	"net/http"
	"sync"
)

type LoadBalancer struct {
	servers []string
	current int
	mu      sync.Mutex
}

func NewLoadBalancer(servers []string) *LoadBalancer {
	return &LoadBalancer{
		servers: servers,
		current: 0,
	}
}

func (lb *LoadBalancer) NextServer() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if len(lb.servers) == 0 {
		return ""
	}

	server := lb.servers[lb.current]
	lb.current = (lb.current + 1) % len(lb.servers)
	return server
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var server string
    cookie, err := r.Cookie("session")
    if err == nil && cookie != nil {
        server = cookie.Value
    }

    if server == "" {
        server = lb.NextServer()
    }

    if server == "" {
        http.Error(w, "No servers available", http.StatusServiceUnavailable)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:  "session",
        Value: server,
        Path:  "/",
    })

    http.Redirect(w, r, "http://"+server+r.RequestURI, http.StatusTemporaryRedirect)
}

func (lb *LoadBalancer) AddServer(server string) {
    lb.mu.Lock()
    defer lb.mu.Unlock()
    lb.servers = append(lb.servers, server)
}

func (lb *LoadBalancer) RemoveServer(server string) {
    lb.mu.Lock()
    defer lb.mu.Unlock()
    for i, s := range lb.servers {
        if s == server {
            lb.servers = append(lb.servers[:i], lb.servers[i+1:]...)
            break
        }
    }
}
