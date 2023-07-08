package di

import (
	"github.com/gin-gonic/gin"
	"github.com/mrtc0/seccamp-2023/pkg/config"
	"github.com/mrtc0/seccamp-2023/pkg/controllers/book"
	"github.com/mrtc0/seccamp-2023/pkg/controllers/healthcheck"
	"github.com/mrtc0/seccamp-2023/pkg/controllers/support/middlewares"
)

func NewServer(
	conf config.Config,
	healthcheckController *healthcheck.HealthcheckController,
	bookController *book.BookController,
) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middlewares.LoggingMiddleware())

	r.GET("/livez", healthcheckController.Livez)
	r.GET("/", bookController.List)

	return r
}
