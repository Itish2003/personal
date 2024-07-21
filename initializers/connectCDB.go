package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var CONTENTDB *gorm.DB

func ConnectCDB() {
	var err error
	dsn := os.Getenv("CONTENTDB")

	CONTENTDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
