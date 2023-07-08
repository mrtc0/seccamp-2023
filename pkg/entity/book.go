package entity

import "context"

type Book struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type BookRepository interface {
	List(ctx context.Context) (*[]Book, error)
}
