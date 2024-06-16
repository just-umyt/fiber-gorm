package items

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/fiber/pkg/db"
	"github.com/just-umyt/fiber/pkg/models"
)

func UpdateItemId(c *fiber.Ctx) error {
	var updateItem models.Item

	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(http.StatusBadRequest).SendString("Invalid ID")
	}

	db.DB.First(&updateItem, id)

	if err := c.BodyParser(&updateItem); err != nil {
		return c.Status(http.StatusBadRequest).SendString("invalid json")
	}

	updateItem.ID = id

	db.DB.Save(&updateItem)

	db.DB.Find(&items)

	return c.JSON(items)
}
