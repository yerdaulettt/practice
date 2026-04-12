package middleware

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type rateLimiter struct {
	visitors map[any]*rate.Limiter
	rate     rate.Limit
	burst    int
	mu       sync.Mutex
}

func NewRateLimiter(r rate.Limit, burst int) *rateLimiter {
	return &rateLimiter{
		visitors: make(map[any]*rate.Limiter),
		rate:     r,
		burst:    burst,
	}
}

func (rl *rateLimiter) getLimiter(id any) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	limiter, exists := rl.visitors[id]
	if !exists {
		limiter = rate.NewLimiter(rl.rate, rl.burst)
		rl.visitors[id] = limiter
	}

	return limiter
}

func (rl *rateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, ok := c.Value("userID").(float64)

		if !ok {
			ip := c.ClientIP()

			if !rl.getLimiter(ip).Allow() {
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests [IP - " + ip + "]"})
				return
			}
		} else {
			if !rl.getLimiter(userId).Allow() {
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "too many requests [User ID - " + strconv.FormatFloat(userId, 'f', -1, 64) + "]"})
				return
			}
		}

		c.Next()
	}
}
