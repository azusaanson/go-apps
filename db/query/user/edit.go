package dbUser

import (
	"context"

	"github.com/azusaanson/invest-api/db"
	"github.com/azusaanson/invest-api/domain"
)

func (uq *userQuery) CreateUser(
	ctx context.Context,
	user domain.User,
) error {
	record := &db.User{
		Name:     string(user.Name()),
		Password: string(user.HashedPassword()),
		Role:     string(user.Role()),
	}

	if err := uq.conn.Create(record).Error; err != nil {
		return err
	}

	return nil
}

func (uq *userQuery) UpdateUser(
	ctx context.Context,
	user domain.User,
) error {
	err := uq.conn.
		Model(&db.User{}).
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

func (uq *userQuery) DeleteUser(
	ctx context.Context,
	userID domain.UserID,
) error {
	err := uq.conn.
		Where("id = ?", userID).
		Delete(&db.User{}).Error
	if err != nil {
		return err
	}

	return nil
}
