package db

import (
	"log"
	"os"

	"github.com/shivaraj-shanthaiah/godoc/docservice/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn, ok := os.LookupEnv("DB_Con")
	if !ok {
		log.Fatal("Error Loading database env")
	}

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connnecting to admin database %v", err.Error())
	}

	DB.AutoMigrate(&models.Doctor{})
	return DB
}
