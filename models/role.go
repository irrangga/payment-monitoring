package models

import (
	"time"

	"gorm.io/gorm"
)

type RoleID struct {
	ID int `json:"role_id"`
}

type Role struct {
	ID        int            `gorm:"primary_key" json:"role_id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	RoleName  string         `json:"role_name"`
}
