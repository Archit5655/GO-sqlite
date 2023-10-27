package database

import (
	models "fiber/Models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstanse struct {
	DB *gorm.DB
}

var Database DBinstanse

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("FAiled to connect to the database")
		os.Exit(2)

	}
	log.Println("COnnected to the database successdfully")
	db.Logger=logger.Default.LogMode(logger.Info)
	log.Println("running mIgrartions")
	//  Add migrations
	db.AutoMigrate(&models.User{}, &models.Products{}, &models.Orders{})

	Database = DBinstanse{DB: db}

}
