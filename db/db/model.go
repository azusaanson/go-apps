package db

import (
	"time"
)

type BaseModel struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
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
