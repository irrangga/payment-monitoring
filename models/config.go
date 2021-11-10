package models

import "gorm.io/gorm"

type ServerConfig struct {
	AppName     string `env:"APP_NAME"`
	AppPort     string `env:"APP_PORT"`
	LogLevel    string `env:"LOG_LEVEL"`
	Environment string `env:"ENVIRONMENT"`
	JWTSecret   string `env:"JWT_SECRET"`
	DBUsername  string `env:"DB_USERNAME"`
	DBPassword  string `env:"DB_PASSWORD"`
	DBHost      string `env:"DB_HOST"`
	DBPort      string `env:"DB_PORT"`
	DBName      string `env:"DB_NAME"`
}

type Auth struct {
	gorm.Model
	AuthUUID   string `gorm:"size:255;not null;" json:"auth_uuid"`
	Username   string `gorm:"not null;" json:"username"`
	RoleID     string `json:"role_id"`
	Name       string `json:"name"`
	CustomerID int    `json:"customer_id" gorm:"-"`
	OfficerID  int    `json:"officer_id" gorm:"-"`
	RoleName   string `json:"role_name"`
}

type JwtModel struct {
	Token      string `json:"token"`
	CustomerID int    `json:"customer_id,omitempty"`
	Name       string `json:"name"`
	OfficerID  int    `json:"officer_id,omitempty"`
}

const (
	Admin          = 1
	GeneralSupport = 2
	Accounting     = 3
	Customer       = 4
)
