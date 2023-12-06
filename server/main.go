package main

import (
	"github.com/anttam/goodOrBadHabitRecorder/server/initializers"
	"github.com/anttam/goodOrBadHabitRecorder/server/router"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.InitializeDb()
	initializers.StartCronJobs()
}
func main (){
		router.Start()
}