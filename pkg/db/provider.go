package db

import (
	"github.com/mrtc0/seccamp-2023/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseProvider struct {
	db *gorm.DB
}

func NewDatabaseProvider(dsn string) *DatabaseProvider {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migration(db)

	return &DatabaseProvider{db: db}
}

func NewDatabaseCon(conf config.Config) *gorm.DB {
	dbp := NewDatabaseProvider(conf.DatabaseConfig.DSN)
	return dbp.db
}
