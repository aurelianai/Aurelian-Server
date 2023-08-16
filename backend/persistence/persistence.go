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
	// TODO Hacky fix for turning on ssl in demo-production
	var SSL_MODE string
	if PG_USER == "e253" {
		SSL_MODE = "verify-full"
	} else {
		SSL_MODE = "disabled"
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", PG_HOST, PG_USER, PG_PASS, PG_DB, PG_PORT, SSL_MODE)

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
