package gomiddlewares

import (
	"github.com/gin-gonic/gin"
)

func SetTraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("X-Trace-Id", c.Request.Header.Get("X-Trace-Id"))
		c.Next()
	}
}