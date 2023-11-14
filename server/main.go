package main

import "github.com/gin-gonic/gin"

func main (){
	request := gin.Default()
	request.GET("/", func(context *gin.Context){
		context.JSON(200, gin.H{
			"message": "hello from the server!",
		})
	})
	request.Run()
}