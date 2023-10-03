package store

import (
	"errors"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormStore struct {
	Client *gorm.DB
}

func NewGorm(dataSourceName string) (store GormStore, err error) {
	if !strings.Contains(dataSourceName, "://") {
		err = errors.New("store: undefined data source name " + dataSourceName)
		return
	}

	store.Client, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return
	}

	sqlDB, err := store.Client.DB()
	if err != nil {
		return
	}
	sqlDB.SetMaxOpenConns(20)

	return
}
