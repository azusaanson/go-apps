package db

import (
	"time"
)

type BaseModel struct {
	ID        uint64
	UpdatedAt time.Time
	CreatedAt time.Time
}

type Session struct {
	BaseModel
	UUID         string
	UserID       uint64
	RefreshToken string
	UserAgent    string
	ClientIp     string
	IsBlocked    bool
	ExpiresAt    time.Time
}

type User struct {
	BaseModel
	Name     string
	Password string
	Role     string
}

type Invest struct {
	BaseModel
	UserID     uint64
	Amount     float64
	Type       string
	InvestedAt time.Time
}
