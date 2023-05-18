package db

import (
	"gorm.io/gorm"
)

type Store struct {
	conn *gorm.DB
	UserQueries
}

type StoreInterface interface {
	UserQueries
}

func NewStore(conn *gorm.DB) StoreInterface {
	return &Store{conn: conn}
}
