package persistence

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Panics on Failure
func InitAndMigrate() {
	PG_HOST := os.Getenv("PG_HOST")
	PG_USER := os.Getenv("PG_USER")
	PG_PASS := os.Getenv("PG_PASS")
	PG_PORT := os.Getenv("PG_PORT")
	PG_DB := os.Getenv("PG_DB")
	var dsn string
	if os.Getenv("GO_ENV") == "prod" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=verify-full", PG_HOST, PG_USER, PG_PASS, PG_DB, PG_PORT)
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", PG_HOST, PG_USER, PG_PASS, PG_DB, PG_PORT)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Error connecting to pg: %s", err.Error()))
	}

	uerr := db.AutoMigrate(&User{})
	cerr := db.AutoMigrate(&Chat{})
	merr := db.AutoMigrate(&Message{})
	if uerr != nil || cerr != nil || merr != nil {
		panic("Error occured during migration")
	}

	DB = db
}
