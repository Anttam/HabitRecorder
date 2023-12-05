package controllers

import (
	"github.com/anttam/goodOrBadHabitRecorder/server/models"

	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateDatabaseEntry(data models.Habit)(error){
	result := DB.Create(data)
	return result.Error
}

func SearchForExistingRecord(requestHabit models.Habit )( models.Habit, error){
	var query models.Habit
	err :=DB.Where(
		"uid = ? AND date = ?",
		 requestHabit.Uid, 
		 requestHabit.Date).First(&query).Error
	return query, err
}