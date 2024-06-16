package items

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/fiber/pkg/db"
	"github.com/just-umyt/fiber/pkg/models"
)

func CreateItem(c *fiber.Ctx) error {
	var newItem models.Item
	if err := c.BodyParser(&newItem); err != nil {
		return c.Status(http.StatusBadRequest).SendString("invalid json")
	}

	newItem.ID = len(items) + 1

	db.DB.Create(&newItem)

	db.DB.Find(&items)

	return c.JSON(items)
}
