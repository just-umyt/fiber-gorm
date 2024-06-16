package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/just-umyt/fiber/pkg/db"
	"github.com/just-umyt/fiber/pkg/items"
)

func main() {
	db.InitDB()

	app := fiber.New()

	app.Get("/items/", items.GetAllItem)
	app.Get("/items/:id", items.GetItemId)
	app.Post("/items/", items.CreateItem)
	app.Put("/items/:id", items.UpdateItemId)
	app.Delete("/items/:id", items.DeleteItemId)

	log.Fatal(app.Listen(":3000"))
}
