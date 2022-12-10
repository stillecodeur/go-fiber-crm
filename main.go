package main

import (
	"go-fiber-crm/database"
	"go-fiber-crm/lead"

	"github.com/gofiber/fiber"
)

func setupRouter(app *fiber.App) {
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.NewLead)
}

func main() {
	app := fiber.New()
	database.Init()
	database.DbConn.AutoMigrate(&lead.Lead{})
	setupRouter(app)
	app.Listen(3000)
	defer database.DbConn.Close()
}
