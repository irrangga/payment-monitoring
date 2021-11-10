package db

import (
	"payment-monitoring/models"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	err := db.Debug().AutoMigrate(
		models.OfficerAccount{},
		models.WorkUnitAccount{},
		models.User{},
		models.Role{},
		models.Auth{},
		models.EntityPayment{},
	)
	if err != nil {
		return
	}
}
