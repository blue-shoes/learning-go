package main

import (
	"fmt"

	"github.com/blue-shoes/go-fiber-crm/database"
	"github.com/blue-shoes/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
	app.Get("/api/v1/lead/:id", lead.GetLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection opened to db")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer closeConnection()
}

func closeConnection() {
	db, _ := database.DBConn.DB()
	db.Close()
}
