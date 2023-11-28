package initializers

import (
"gorm.io/driver/postgres"
"gorm.io/gorm"
"os"
"log"

)

func ConnectToDb () (*gorm.DB){
	signOn := os.Getenv("DATABASEURL")
	db, err := gorm.Open(postgres.Open(signOn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
	log.Print("Connected to DB")
	return db
}