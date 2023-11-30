package router

import (

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/anttam/goodOrBadHabitRecorder/server/controllers"
)
type RequestBody struct{
	Habit string 
}


func Start() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Origin","Content-Type"},
	}))

	r.POST("/", controllers.HandlePost)

	r.Run(":8000")
}