package library

import "empty/internal/domain/book"

type Service struct {
	bookRepository book.Repository
}
