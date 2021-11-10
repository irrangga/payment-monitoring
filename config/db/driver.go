package db

import (
	"payment-monitoring/config/env"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	var configDb = env.Config

	if DB != nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("recovery error")
		}
	}()

	urlDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", configDb.DBUsername, configDb.DBPassword, configDb.DBHost, configDb.DBPort, configDb.DBName)

	DB, err = gorm.Open(mysql.Open(urlDB), &gorm.Config{})
	if err != nil {
		log.Println("cannot connect to DB")
	}

	log.Print("success connect to db")

	AutoMigrate(DB)

}
