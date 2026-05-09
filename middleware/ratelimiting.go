package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	Limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)
)

func getClientIP(c *gin.Context) string {
	ip := c.ClientIP()
	// when use proxy, the ClientIP() may return empty string, so we can get the IP from RemoteAddr
	if ip == "" {
		ip = c.Request.RemoteAddr
	}

	return ip
}

func getRateLimitKey(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	client, exists := clients[ip]
	if !exists {
		limiter := rate.NewLimiter(5, 10) // 5 requests per second with a burst of 10
		clients[ip] = &Client{limiter, time.Now()}

		return limiter
	}

	client.lastSeen = time.Now()
	return client.Limiter
}

func CleanupClients() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

// Test:  ab -n 20 -c 1 -H "X-API-Key:api_key" localhost:8081/api/v1/user
func RateLimitingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := getClientIP(c)

		limiter := getRateLimitKey(ip)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":   "Too many requests",
				"message": "You have exceeded the rate limit. Please try again later.",
			})
		}

		c.Next()
	}
}
