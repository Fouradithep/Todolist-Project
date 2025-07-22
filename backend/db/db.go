package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fouradithep/todolist/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Init() {

	// โหลดค่า .env
	err := godotenv.Load()
	if err != nil {
		log.Println("ไม่สามารถโหลดไฟล์ .env:", err)
	}

	// อ่านค่าจาก environment
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)
		fmt.Println("DSN:", dsn)


	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal("ไม่สามารถเชื่อมต่อฐานข้อมูล:", err)
	}

	// สร้างตาราง tasks อัตโนมัติ
	err = db.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		log.Fatal("AutoMigrate ล้มเหลว:", err)
	}

	DB = db
	fmt.Println("เชื่อมต่อฐานข้อมูลสำเร็จด้วย GORM")
}
