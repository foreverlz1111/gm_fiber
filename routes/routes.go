package routes

import (

	"gorm-mysql/database"
	"gorm-mysql/models"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
    student := new(models.Student)
    result := database.DB.Find(&student)
	return c.SendString("hi~,found :"+ string(result.RowsAffected))
}

