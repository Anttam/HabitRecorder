package models


type Habit struct {
	ID string `gorm:"primary key"`
	Uid string
	Date string
	HabitCount int
}