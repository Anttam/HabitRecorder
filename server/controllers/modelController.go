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

func SearchForExistingRecord(uid string, date string )( models.Habit, error){
	var query models.Habit
	err :=DB.Where(
		"uid = ? AND date = ?",
		 uid, 
		 date).First(&query).Error
	return query, err
}