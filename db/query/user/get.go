package user

import (
	"context"

	"errors"

	"github.com/azusaanson/invest-api/db/model"
	"github.com/azusaanson/invest-api/domain"
	"gorm.io/gorm"
)

type UserQuery struct{ conn *gorm.DB }

func NewUserQuery(conn *gorm.DB) domain.UserQueries {
	return &UserQuery{conn: conn}
}

func (uq *UserQuery) GetUserByName(
	ctx context.Context,
	name domain.UserName,
) (*domain.User, error) {
	record := &model.User{}

	err := uq.conn.Model(&model.User{}).
		Where("name = ?", name).
		First(record).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if record.ID == 0 {
		return nil, nil
	}

	return domain.NewUserFromSource(record.ID, record.Name, record.Password, record.Role)
}
