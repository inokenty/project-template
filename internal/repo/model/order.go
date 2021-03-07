package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	UserID    uint
	User      *User
	CreatedAt *time.Time
	Sum       int64
}
