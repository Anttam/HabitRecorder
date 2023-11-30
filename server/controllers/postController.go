package controllers

import (
	"errors"
	"time"

	"github.com/anttam/goodOrBadHabitRecorder/server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type RequestBody struct{
	Habit string 
}

var DB *gorm.DB

func HandlePost (c *gin.Context){
		var request RequestBody
		requestValue := 0;
		 if err := c.BindJSON(&request); 
		 err!= nil{
			panic(err)
		 }
		 // assign 1 for good -1 for bad
		 if request.Habit == "good"{
			requestValue++
		 }else if request.Habit == "bad"{
			requestValue--
		 }else{
			c.Data(400, "application/json",[]byte(`{"message":"Invalid button type"}`))
		 }

		 //declare habit and query
		 habit := models.Habit{
			ID:uuid.New().String(),
			Uid: "test",
			Date: time.Now().Format("01-02-2006"),
			HabitCount: requestValue,}

			var query models.Habit

			// check if record exists for user ID and date
		   err:= DB.Where("uid = ? AND date = ?", habit.Uid, habit.Date).First(&query).Error
			 if err != nil{
				//if no record, create one
			 	if errors.Is(err, gorm.ErrRecordNotFound){
					err := createDatabaseEntry(habit)
					if err!= nil{
						c.Data(400, "application/json", []byte(`{"message" : "unable to create record"}`))
					}else{
					c.Data(201, "application/json", []byte(`{"message" : "No data entry, created new entry"}`))
					}
				}else{
				//else send error
				c.Data(400, "application/json", []byte(`{"message" : "unable to create record"}`))
				}
			}
		 // if, update record with habit 
		 if query.ID != ""{
			query.HabitCount = query.HabitCount + requestValue;
			DB.Save(&query)
		 }

		 // send success message
}

func createDatabaseEntry(data models.Habit)(error){
	result := DB.Create(data)
	return result.Error
}

