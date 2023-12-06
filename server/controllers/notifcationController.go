package controllers

import (
	"fmt"
	"net/smtp"
	"os"
	"time"

	"gorm.io/gorm"
)

func SendEmailNotifcation(){
	auth := smtp.PlainAuth(
		"Habit Tracker",
		os.Getenv("EMAIL"),
		os.Getenv("EMAILPASSWORD"),
		"smtp.gmail.com",
	)
	var body string
		habit, err := SearchForExistingRecord("test", time.Now().Format("01-02-2006"))
		if err == gorm.ErrRecordNotFound{
			body = "you didn't record a habit for today"
		}else{
			panic(err)
		}

	subject := "Subject: Your Daily Habit Count \n"
	if habit.HabitCount >= 1{
		body = "good job! you had a good habit day!"
	}else if habit.HabitCount < 0{
		body = "you didn't have a great habit day, don't let this set back get you down!"
	}else{
		body = "you had a neutral habit day"
	}

	msg:= subject+body
	

	emailerr:= smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"Habit Tracker",
		[]string{"anttam95@gmail.com"},
		[]byte(msg),
	)

	if emailerr != nil{
		panic(emailerr)
	}else{
		fmt.Printf("email sent for %s", time.Now().Format("01-02-2006"))
	}
}