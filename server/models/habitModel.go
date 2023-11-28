package models

import "time"

type Habit struct {
	ID uint `gorm: "primary key"`
	Uid uint
	Date time.Time
	HabitCount int
}