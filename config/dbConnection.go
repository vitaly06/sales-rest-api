package config

import (
	"fmt"
	"os"

	"github.com/vitaly06/sales-rest-api/Models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", db_host, db_user, db_password, db_name, db_port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("db connection failed")
	}

	DB = db

	fmt.Println("db connected successfully")

	AutoMigrate(db)
}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&Models.Cashier{},
		&Models.Category{},
		&Models.Discount{},
		&Models.Order{},
		&Models.Payment{},
		&Models.PaymentType{},
		&Models.Product{},
	)
}
