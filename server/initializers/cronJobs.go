package initializers

import (
	"github.com/anttam/goodOrBadHabitRecorder/server/controllers"

	"github.com/robfig/cron/v3"

	"log"
)

func StartCronJobs () {
	cr := cron.New()
	//declare CRON job to send email notifaction at 9 PM
	cr.AddFunc("0 21 * * *", controllers.SendEmailNotifcation)
		//start CRON job
		cr.Start()
		log.Print("CRON job started")
}