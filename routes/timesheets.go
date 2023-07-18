package routes

import (
	"errors"

	"github.com/Ymit24/fiber-api/database"
	"github.com/Ymit24/fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type TimeEntryDTO struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Start string `json:"start"`
	Stop  string `json:"stop"`
}

type TimeSheetDTO struct {
	ID      uint           `json:"id"`
	Title   string         `json:"title"`
	Entries []TimeEntryDTO `json:"entries"`
}

func createTimeEntryDto(entry models.TimeEntry) TimeEntryDTO {
	return TimeEntryDTO{
		ID:    entry.ID,
		Name:  entry.Name,
		Start: entry.Start,
		Stop:  entry.Stop,
	}
}

func createTimeSheetDto(timesheet models.TimeSheet, entries []models.TimeEntry) TimeSheetDTO {
	var entryDtos []TimeEntryDTO
	for _, entry := range entries {
		entryDto := createTimeEntryDto(entry)
		entryDtos = append(entryDtos, entryDto)
	}

	return TimeSheetDTO{
		ID:      timesheet.ID,
		Title:   timesheet.Title,
		Entries: entryDtos,
	}
}

func findTimeSheet(id int, timesheet *models.TimeSheet) error {
	database.Database.Db.Find(&timesheet, "id = ?", id)
	if timesheet.ID == 0 {
		return errors.New("timesheet does not exist")
	}
	return nil
}

func findTimeEntryForTimeSheet(id int, entries *[]models.TimeEntry) error {
	// NOTE: write this
	return nil
}

func CreateTimeSheet(c *fiber.Ctx) error {
	var timesheet models.TimeSheet

	if err := c.BodyParser(&timesheet); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var entries []models.TimeEntry

	timesheetDto := createTimeSheetDto(timesheet, entries)

	return c.Status(200).JSON(timesheetDto)
}

// func DeleteTimeSheet(c *fiber.Ctx) error {}
// func UpdateTimeSheet(c *fiber.Ctx) error {}

func GetTimeSheet(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var timesheet models.TimeSheet

	if err := findTimeSheet(id, &timesheet); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var entries []models.TimeEntry

	if err := findTimeEntryForTimeSheet(id, &entries); err != nil {
	}

	timesheetDto := createTimeSheetDto(timesheet, entries)
	return c.Status(200).JSON(timesheetDto)
}

// func GetTimeSheets(c *fiber.Ctx) error   {}
