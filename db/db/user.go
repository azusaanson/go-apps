package db

import (
	"context"

	"errors"

	"github.com/azusaanson/invest-api/domain"
	"gorm.io/gorm"
)

type UserQueries interface {
	GetUserByName(ctx context.Context, name domain.UserName) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, userID domain.UserID) error
}

func (s *Store) GetUserByName(
	ctx context.Context,
	name domain.UserName,
) (*domain.User, error) {
	record := &User{}

	err := s.conn.Model(&User{}).
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

func (s *Store) CreateUser(
	ctx context.Context,
	user *domain.User,
) error {
	record := &User{
		Name:     string(user.Name()),
		Password: string(user.HashedPassword()),
		Role:     string(user.Role()),
	}

	if err := s.conn.Create(record).Error; err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateUser(
	ctx context.Context,
	user *domain.User,
) error {
	err := s.conn.
		Model(&User{}).
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

func (s *Store) DeleteUser(
	ctx context.Context,
	userID domain.UserID,
) error {
	err := s.conn.
		Where("id = ?", userID).
		Delete(&User{}).Error
	if err != nil {
		return err
	}

	return nil
}
