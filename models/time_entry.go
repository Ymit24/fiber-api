package models

import "time"

type TimeEntry struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	TimesheetId int       `json:"timesheet_id"`
	Timesheet   TimeSheet `gorm:"foreignKey:TimesheetId"`
	Name        string    `json:"name"`
	Start       string    `json:"start"`
	Stop        string    `json:"stop"`
	CreatedAt   time.Time
}
