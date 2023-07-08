package db

import (
	"context"

	"github.com/mrtc0/seccamp-2023/pkg/entity"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) entity.BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) List(ctx context.Context) (*[]entity.Book, error) {
	var books []entity.Book
	result := r.db.WithContext(ctx).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return &books, nil
}
