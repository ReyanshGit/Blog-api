package config

import (
	"fmt"
	"os"
	
	"blogapi/models"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// .env se values lo
	user     := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host     := os.Getenv("DB_HOST")
	port     := os.Getenv("DB_PORT")
	dbname   := os.Getenv("DB_NAME")

	// Connection string
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		user, password, host, port, dbname,
	)

	// Connect karo
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Database connect nahi hua: " + err.Error())
	}

	// Tables automatically ban jayenge
	db.AutoMigrate(&models.User{}, &models.Post{})

	DB = db
	fmt.Println("✅ Database connected!")
}