package database

import (
	"fmt"
	"os"

	"github.com/Shalqarov/student-crud/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("postgres://%s:%s@%s%s/%s?sslmode=disable",
		user, password, host, port, name,
	)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&repository.Student{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
