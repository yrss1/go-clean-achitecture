package postgres

import (
	"context"
	"empty/internal/domain/book"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}
func (r *BookRepository) List(ctx context.Context) (dest []book.Entity, err error) {
	err = r.db.WithContext(ctx).Find(&dest).Error
	return
}

func (r *BookRepository) Get(ctx context.Context) (dest book.Entity, err error) {
	err = r.db.WithContext(ctx).First(&dest).Error
	return
}

func (r *BookRepository) Add(ctx context.Context, data book.Entity) (id string, err error) {
	err = r.db.WithContext(ctx).Create(&data).Scan(&id).Error
	return
}

//func (r *BookRepository) Update(ctx context.Context) {
//
//}

func (r *BookRepository) Delete(ctx context.Context, id string) (err error) {
	err = r.db.Delete(&id).Error
	return
}
