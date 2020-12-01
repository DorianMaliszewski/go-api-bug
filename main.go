package main

import (
	"bug-tracker/backend/config"
	"bug-tracker/backend/daos"
	"bug-tracker/backend/handlers"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New()

	app.Use(logger.New())

	// Root API route
	api := app.Group("/api") // /api

	// API v1 routes
	v1 := api.Group("/v1") // /api/v1
	bugsRouter := v1.Group("/bugs")
	handlers.SetupBugRoutes(bugsRouter)

	c := config.GetConfiguration()

	daos.GetConn()
	log.Fatal(app.Listen(fmt.Sprintf(":%s", c.Port)))
}
