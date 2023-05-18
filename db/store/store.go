package store

import (
	"github.com/azusaanson/invest-api/db/query/user"
	"github.com/azusaanson/invest-api/domain"
	"gorm.io/gorm"
)

type Store struct {
	conn      *gorm.DB
	UserQuery domain.UserQueries
}

func NewStore(conn *gorm.DB) *Store {
	return &Store{
		conn:      conn,
		UserQuery: user.NewUserQuery(conn),
	}
}
