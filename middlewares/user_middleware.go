package middlewares

import (
    "log"
    "time"
    "github.com/gin-gonic/gin"
)

// LoggingMiddleware logs incoming requests
func LoggingMiddleware() gin.HandlerFunc {

    return func(c *gin.Context) {

        start := time.Now()
        c.Next()
        duration := time.Since(start)
        log.Printf("[LOGGER] Request: %s %s | Duration: %v", c.Request.Method, c.Request.URL.Path, duration)

    }

}
