package library

import (
	"context"
	"empty/internal/domain/book"
)

func (s *Service) ListBooks(ctx context.Context) (res []book.Response, err error) {
	data, err := s.bookRepository.List(ctx)
	if err != nil {
		panic("failed to select")
		return
	}
	res = book.ParseFromEntities(data)

	return
}
func (s *Service) GetBook(ctx context.Context, id string) (res book.Response, err error) {
	data, err := s.bookRepository.Get(ctx, id)
	if err != nil {
		panic("failed to get by id")
		return
	}
	res = book.ParseFromEntity(data)
	return
}

func (s *Service) CreateBook(ctx context.Context, req book.Request) (res book.Response, err error) {
	data := book.Entity{
		Name:  &req.Name,
		Genre: &req.Genre,
		ISBN:  &req.ISBN,
	}
	data.ID, err = s.bookRepository.Add(ctx, data)
	if err != nil {
		panic("failed to create")
		return
	}
	res = book.ParseFromEntity(data)

	return
}

func (s *Service) DeleteBook(ctx context.Context, id string) (err error) {
	err = s.bookRepository.Delete(ctx, id)
	if err != nil {
		panic("failed to delete by id")
		return
	}
	return
}
