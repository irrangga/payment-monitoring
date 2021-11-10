package env

import (
	"log"
	"payment-monitoring/models"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var Config models.ServerConfig

func init() {
	err := loadConfig()
	if err != nil {
		log.Println("cannot open file .env.local")
	}
}

func loadConfig() error {
	err := godotenv.Load()
	if err != nil {
		logrus.Warning(err, "config/env: load file env")
	}

	err = env.Parse(&Config)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return err
}
