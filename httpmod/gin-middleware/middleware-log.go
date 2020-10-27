package gin_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/peywong-galaxy/s-gw/log"
	"go.uber.org/zap"
	"time"
)

func AddLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		reqAddress := c.Request.RemoteAddr
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		log.Logger.Info("gin start to process request",
			zap.String("Method",reqMethod),
			zap.String("Uri", reqUri),
			zap.String("Address", reqAddress),
			zap.Int("Status", statusCode),
			zap.String("client",clientIP))

		c.Next()
		endTime := time.Now()
		Duration := endTime.Sub(startTime)

		log.Logger.Info("gin end in process request",
			zap.Duration("Duration", Duration))
	}
}
