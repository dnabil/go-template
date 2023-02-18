package middleware

import (
	"go-template/sdk/apires"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func Timeout(t time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(t),
		timeout.WithHandler(func(c *gin.Context) {
		  c.Next()
		}),
		timeout.WithResponse(func(c *gin.Context) {
			apires.FailOrError(c, 504, "server timeout", nil)
		}),
	  )
}