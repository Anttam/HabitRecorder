package initializers

import (
	"github.com/anttam/goodOrBadHabitRecorder/server/controllers"
	"github.com/anttam/goodOrBadHabitRecorder/server/models"

"gorm.io/driver/postgres"
"gorm.io/gorm"

"os"
"log"
)
func InitializeDb (){
	db := connectToDb()
	migrateDB(db)
	declareControllerDB(db)
	log.Print("Database initilization completed")
}


func connectToDb () (*gorm.DB){
	signOn := os.Getenv("DATABASEURL")
	db, err := gorm.Open(postgres.Open(signOn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
	log.Print("Connected to DB")

	return db
}

func declareControllerDB (db *gorm.DB) {
	controllers.DB = db
log.Print("Set controllers DB")
}

func migrateDB (db *gorm.DB) {
	db.AutoMigrate(&models.Habit{})
}
