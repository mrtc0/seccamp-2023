package db

import (
	"github.com/mrtc0/seccamp-2023/pkg/entity"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&entity.Book{})
}
