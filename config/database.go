package config

import (
	"fmt"
	"os"

	md "ems-aquadev/models"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	dsn struct {
		Host string
		User string
		Password string
		Dbname string
		Port string
		Timezone string
	}
)

var (
	DB *gorm.DB
	err error
)

func Database() {
	dsn := dsn{
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Dbname: os.Getenv("DB_NAME"),
		Port: os.Getenv("DB_PORT"),
		Timezone: os.Getenv("DB_TIMEZONE"),
	}
	db_url := "user="+dsn.User+" password="+dsn.Password+" dbname="+dsn.Dbname+" port="+dsn.Port+" TimeZone="+dsn.Timezone
	DB, err = gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected")
}

func Migrate() {
	DB.AutoMigrate(
		&md.User{},
		&md.UserProfile{},
		&md.UserAddress{},
		&md.UserPayment{},
		&md.Product{},
		&md.ProductCategory{},
		&md.CartSession{},
		&md.CartItem{},
		&md.Order{},
		&md.OrderItem{},
		&md.PaymentDetails{},
		&md.Admin{},
	)
}