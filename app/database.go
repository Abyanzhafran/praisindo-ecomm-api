package app

import (
	"dev/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSERNAME")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGNAME")
	port := os.Getenv("PGPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	fmt.Println("dsn logging : ", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// use "models" instead of "domain"
	// if err := db.AutoMigrate(&domain.User{}, &domain.Photo{}, &domain.Comment{}, &domain.SocialMedia{}); err != nil {
	// 	log.Fatal(err.Error())
	// }

	if err := db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatal(err.Error())
	}

	return db
}
