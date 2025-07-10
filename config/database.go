package config

import (
	"fmt"
	"log"
	"myapp/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// .env faylni yuklash
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// .env dagi ma'lumotlarni o'qish
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")

	// DSN yaratish
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	// GORM orqali ulanish
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	database.AutoMigrate(
		&models.Category{}, 
		&models.Product{}, 
		&models.Order{}, 
		&models.PermissionGroup{}, 
		&models.Permission{}, 
		&models.Role{}, 
		&models.RolePermission{},
	)
	DB = database
	fmt.Println("✅ PostgreSQL bazaga muvaffaqiyatli ulandi!")
}
