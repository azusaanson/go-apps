package dbUser

import (
	"context"

	"errors"

	"github.com/azusaanson/invest-api/db"
	"github.com/azusaanson/invest-api/domain"
	"gorm.io/gorm"
)

type userQuery struct{ conn *gorm.DB }

func NewUserQuery(conn *gorm.DB) domain.UserQuery {
	return &userQuery{conn: conn}
}

func (uq *userQuery) GetUserByName(
	ctx context.Context,
	name domain.UserName,
) (*domain.User, error) {
	record := &db.User{}

	err := uq.conn.Model(&db.User{}).
		Where("name = ?", name).
		First(record).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if record.ID == 0 {
		return nil, nil
	}

	return domain.NewUser(record.ID, record.Name, record.Password, record.Role)
}
