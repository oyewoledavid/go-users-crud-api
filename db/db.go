package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Course    string `json:"course"`
	Track     string `json:"track"`
	Status    string `json:"status"`
}

func InitDB() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Auto migrate
	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Seed default users if none exist
	var count int64
	DB.Model(&User{}).Count(&count)
	if count == 0 {
		defaultUsers := []User{
			{ID: "01", Firstname: "Matthew", Lastname: "Ojo", Course: "Software Engineering", Track: "FE", Status: "Student"},
			{ID: "02", Firstname: "Moses", Lastname: "Simon", Course: "Software Engineering", Track: "BE", Status: "Student"},
			{ID: "03", Firstname: "Peter", Lastname: "Emeka", Course: "Product", Track: "UI/UX", Status: "Graduated"},
			{ID: "04", Firstname: "David", Lastname: "Femi", Course: "Software Engineering", Track: "DO", Status: "Graduated"},
			{ID: "05", Firstname: "Kingsley", Lastname: "Ebuka", Course: "Data", Track: "DE", Status: "Student"},
		}
		DB.Create(&defaultUsers)
	}
}
