package db

import (
	"fmt"
	"log"

	"github.com/cavalheirodev/finance-app-bff/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var session *gorm.DB

func Initialize() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT)
	instance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	session = instance
}

func AutoMigrate(models ...any) {
	session.AutoMigrate(models...)
}
