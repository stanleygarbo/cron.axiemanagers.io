package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	dsn := "host=127.0.0.1 port=5432 user=postgres dbname=axiemanagers sslmode=disable password=haynako"
	DBConn, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("Error connecting to DB: %v\n", err)
	}
	fmt.Println("Connected to DB sucessfully!")
}
