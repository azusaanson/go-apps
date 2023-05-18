package user

import (
	"context"

	"github.com/azusaanson/invest-api/db/model"
	"github.com/azusaanson/invest-api/domain"
)

func (uq *UserQuery) CreateUser(
	ctx context.Context,
	user *domain.User,
) error {
	record := &model.User{
		Name:     string(user.Name()),
		Password: string(user.HashedPassword()),
		Role:     string(user.Role()),
	}

	if err := uq.conn.Create(record).Error; err != nil {
		return err
	}

	return nil
}

func (uq *UserQuery) UpdateUser(
	ctx context.Context,
	user *domain.User,
) error {
	err := uq.conn.
		Model(&model.User{}).
		Where("id = ?", user.ID()).
		Updates(map[string]interface{}{
			"name":     user.Name(),
			"password": user.HashedPassword(),
			"role":     user.Role(),
		}).Error
	if err != nil {
		return err
	}

	return nil
}

func (uq *UserQuery) DeleteUser(
	ctx context.Context,
	userID domain.UserID,
) error {
	err := uq.conn.
		Where("id = ?", userID).
		Delete(&model.User{}).Error
	if err != nil {
		return err
	}

	return nil
}
