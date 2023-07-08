package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		logrus.WithFields(logrus.Fields{
			"method":    c.Request.Method,
			"uri":       c.Request.RequestURI,
			"status":    c.Writer.Status(),
			"client_ip": c.ClientIP(),
		}).Info("request completed")

		c.Next()
	}
}
