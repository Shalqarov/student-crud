package main

import (
	"log"

	"github.com/Shalqarov/student-crud/api"
	"github.com/Shalqarov/student-crud/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	app := api.Setup(db)
	app.Listen(":8080")
}
