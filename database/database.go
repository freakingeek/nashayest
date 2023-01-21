package database

import (
	"log"
	"os"

	"github.com/sttatusx/nashayest/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatal("Database connection failed!\n", err.Error())
		os.Exit(2)
	}

	log.Println("Database connected successfully!")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations ...")
	db.AutoMigrate(&models.Word{})

	Database = db

	return db
}
