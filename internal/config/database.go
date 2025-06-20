package database_config

import (
	"expense-tracker/api/models"
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		name     = os.Getenv("DB_NAME")
	)

	db_port, err := strconv.Atoi(port)

	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, db_port, user, name, password)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	err = db.AutoMigrate(
		models.Users{},
		models.Categories{},
		models.Transactions{},
		models.Budgets{},
	)

	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}

	fmt.Println("Database connected")

	return db
}
