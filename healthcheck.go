package main

import(
	"net/http"
	"time"
)

func CheckHealth(server string) bool{
	resp, err := http.Get("http://"+server+"/health")
	if err != nil{
		return false
	}
	return resp.StatusCode == http.StatusOK
}

func (lb *LoadBalancer) PeriodicHealthCheck(interval time.Duration){
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C{
		lb.mu.Lock()
		for i := len(lb.servers) - 1; i>=0; i--{
			if !CheckHealth(lb.servers[i]){
				lb.servers = append(lb.servers[:i], lb.servers[i+1:]...)
			}
		}
		lb.mu.Unlock()
	}
}