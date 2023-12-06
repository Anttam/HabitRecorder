package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/anttam/goodOrBadHabitRecorder/server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type RequestBody struct{
	Habit string 
	Uid string
}
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
			Uid: request.Uid,
			Date: time.Now().Format("01-02-2006"),
			HabitCount: requestValue,}

	
			// check if record exists for user ID and date
		   query,err:= SearchForExistingRecord(habit.Uid, habit.Date)
			 if err != nil{
				//if no record, create one
			 	if errors.Is(err, gorm.ErrRecordNotFound){
					err := CreateDatabaseEntry(habit)
						if err!= nil{
							//handle error creating entry
							c.JSON(
								400,
								gin.H{"message" : "unable to create record"})
						}else{
						c.Status(204)
						}
			//else send error for other than No Record error
				}else{
				c.JSON(
					400,
					gin.H{"message" : "unable to create record"})
			}
		}
		 // if record found, update record with habit 
		 fmt.Println(query.ID)
		 if query.ID != ""{
			query.HabitCount = query.HabitCount + requestValue;
			DB.Save(&query)
			 // send success message
			 c.JSON(201, gin.H{"message":"record found. Count updated"})
			// c.Data(204, "application/json", []byte(`{"message": "record found. Count updated"}`))
		 }

		
}
