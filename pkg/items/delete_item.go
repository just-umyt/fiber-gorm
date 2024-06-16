package items

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/fiber/pkg/db"
	"github.com/just-umyt/fiber/pkg/models"
)

func DeleteItemId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(http.StatusBadRequest).SendString("Invalid ID")
	}

	db.DB.Delete(&models.Item{}, id)

	db.DB.Find(&items)

	return c.JSON(items)
}
