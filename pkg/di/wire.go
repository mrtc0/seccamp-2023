//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/mrtc0/seccamp-2023/pkg/config"
	"github.com/mrtc0/seccamp-2023/pkg/controllers/book"
	"github.com/mrtc0/seccamp-2023/pkg/controllers/healthcheck"
	"github.com/mrtc0/seccamp-2023/pkg/db"
	"github.com/mrtc0/seccamp-2023/pkg/usecases"
)

var APISet = wire.NewSet(
	healthcheck.NewHealthcheckController,
	book.NewBookController,
)

var DBSet = wire.NewSet(
	db.NewDatabaseCon,
	db.NewBookRepository,
)

var UsecaseSet = wire.NewSet(
	usecases.NewListBooks,
)

func Initialize(conf config.Config) *gin.Engine {
	wire.Build(
		UsecaseSet,
		NewServer,
		APISet,
		DBSet,
	)

	return nil
}
