package repository

import (
	"empty/internal/domain/book"
	"empty/pkg/store"
)

type Configuration func(r *Repository) error

type Repository struct {
	postgres store.GormStore

	Book book.Repository
}

//func New() {
//
//}
