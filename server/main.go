package main

import (
	"github.com/anttam/goodOrBadHabitRecorder/server/router"
	"github.com/joho/godotenv"
	"github.com/anttam/goodOrBadHabitRecorder/server/models"
	"github.com/anttam/goodOrBadHabitRecorder/server/initializers"
	// "fmt"
	"gorm.io/gorm"
)

var DB *gorm.DB
func init(){
	godotenv.Load()
	DB = initializers.ConnectToDb()
	DB.AutoMigrate(&models.Habit{})
}

func main (){
		router.Start()
}