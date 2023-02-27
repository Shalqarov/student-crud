package database

import (
	"fmt"
	"os"

	"github.com/Shalqarov/student-crud/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect takes the data from the env file and then connects to the database,
// returns the connection and the error (if there is one)
func Connect() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname,
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
