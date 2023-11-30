package main

import (
	"github.com/anttam/goodOrBadHabitRecorder/server/controllers"
	"github.com/anttam/goodOrBadHabitRecorder/server/initializers"
	"github.com/anttam/goodOrBadHabitRecorder/server/models"
	"github.com/anttam/goodOrBadHabitRecorder/server/router"

	"github.com/joho/godotenv"
)


func init(){
	godotenv.Load()
	DB := initializers.ConnectToDb()
	DB.AutoMigrate(&models.Habit{})
	controllers.DB = DB
}

func main (){
	
		router.Start()
	
}