package items

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/fiber/pkg/db"
	"github.com/just-umyt/fiber/pkg/models"
)

func GetItemId(c *fiber.Ctx) error {
	var item models.Item

	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(http.StatusBadRequest).SendString("Invalid ID")
	}

	db.DB.First(&item, id)

	return c.JSON(item)

}
