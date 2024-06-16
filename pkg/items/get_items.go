package items

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/fiber/pkg/db"
	"github.com/just-umyt/fiber/pkg/models"
)

var items []models.Item

func GetAllItem(c *fiber.Ctx) error {

	result := db.DB.Find(&items)

	if result.Error != nil {
		c.Status(http.StatusNotFound).SendString("cant get items")
	}

	return c.JSON(items)
}
