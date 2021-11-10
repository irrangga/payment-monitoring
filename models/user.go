package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID int `gorm:"primary_key" json:"user_id"`
}

type OfficerAccount struct {
	ID        int            `gorm:"primary_key" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    int            `json:"user_id"`
	Name      string         `json:"name"`
	LoginAs   string         `json:"login_as,omitempty" gorm:"-"`
	Role      int            `json:"role"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
}

type WorkUnitAccount struct {
	ID        int            `gorm:"primary_key" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    int            `json:"user_id"`
	Name      string         `json:"name"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
}
