package router

import (
	"fmt"
	

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	r.POST("/", func(c *gin.Context){
		var request RequestBody
		 if err := c.BindJSON(&request); err!= nil{
			panic(err)
		 }
		 fmt.Println(request.Habit)

	})

	r.Run(":8000")
}