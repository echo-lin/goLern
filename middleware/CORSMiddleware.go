package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-All-Origin", "http://127.0.0.1:6071,http://127.0.0.1:8081")
		c.Writer.Header().Set("Access-Control-Max-Age", "80400")
		c.Writer.Header().Set("Access-Control-All-Methods", "*")
		c.Writer.Header().Set("Access-Control-All-Headers", "*")
		c.Writer.Header().Set("Access-Control-All-Credentials", "*")
		if c.Request.Method == http.MethodOptions {
			fmt.Println("url:", c.Request.RequestURI)
		} else {
			c.Next()
		}
	}
}
