package domain

import (
	"context"
	"errors"
	"fmt"
)

type User struct {
	id             UserID
	name           UserName
	hashedPassword HashedPassword
	role           UserRole
}

func (u *User) ID() UserID                     { return u.id }
func (u *User) Name() UserName                 { return u.name }
func (u *User) HashedPassword() HashedPassword { return u.hashedPassword }
func (u *User) Role() UserRole                 { return u.role }

func NewUser(
	id uint64,
	name string,
	hashedPassword string,
	role string,
) (*User, error) {
	newID, err := NewUserID(id)
	if err != nil {
		return nil, err
	}

	newName, err := NewUserName(name)
	if err != nil {
		return nil, err
	}

	newHashedPassword, err := NewHashedPassword(hashedPassword)
	if err != nil {
		return nil, err
	}

	newRole, err := NewUserRole(role)
	if err != nil {
		return nil, err
	}

	return &User{
		id:             newID,
		name:           newName,
		hashedPassword: newHashedPassword,
		role:           newRole,
	}, nil
}

type UserID uint64

var ErrUserIDZero = errors.New("user id: must not be zero")

func NewUserID(v uint64) (UserID, error) {
	if v == 0 {
		return 0, ErrUserIDZero
	}

	return UserID(v), nil
}

type UserName string

const UserNameMaxLength = 32

var (
	ErrUserNameEmpty   = errors.New("user name: must not be empty")
	ErrUserNameTooLong = fmt.Errorf("user name: must not be longer than %d characters", UserNameMaxLength)
)

func NewUserName(v string) (UserName, error) {
	if v == "" {
		return "", ErrUserNameEmpty
	}

	if len([]rune(v)) > UserNameMaxLength {
		return "", ErrUserNameTooLong
	}

	return UserName(v), nil
}

type HashedPassword string

var ErrHashedPasswordEmpty = errors.New("hashed password: must not be empty")

func NewHashedPassword(v string) (HashedPassword, error) {
	if v == "" {
		return HashedPassword(""), ErrHashedPasswordEmpty
	}

	return HashedPassword(v), nil
}

type UserRole string

const (
	RoleUser  UserRole = "user"
	RoleAdmin UserRole = "admin"
)

func NewUserRole(v string) (UserRole, error) {
	switch v {
	case string(RoleUser):
		return RoleUser, nil
	case string(RoleAdmin):
		return RoleAdmin, nil
	}

	return UserRole(""), nil
}

type UserQuery interface {
	GetUserByName(ctx context.Context, name UserName) (*User, error)
	CreateUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, user User) error
	DeleteUser(ctx context.Context, userID UserID) error
}
