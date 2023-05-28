package db

import (
	"context"

	"gorm.io/gorm"
)

type Store struct {
	conn *gorm.DB
}

type StoreInterface interface {
	ExecTx(ctx context.Context, fn func(context.Context) error) error
	UserQueries
}

func NewStore(conn *gorm.DB) StoreInterface {
	return &Store{conn: conn}
}

func (s *Store) ExecTx(ctx context.Context, fn func(context.Context) error) error {
	tx := s.conn.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer tx.Rollback()

	if err := fn(ctx); err != nil {
		return err
	}

	return tx.Commit().Error
}
