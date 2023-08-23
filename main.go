package main

import (
	"gorm-preload/database"
	"gorm-preload/middleware"
	"gorm-preload/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// create app
	app := fiber.New()

	// load env
	godotenv.Load()

	// use cors
	app.Use(cors.New())

	// use middleware
	app.Use(middleware.CustomeMiddleware())

	// setup database
	database.DBConn = database.InitDb()

	// setup routes
	routes.RoutesInit(app)

	// set port
	port, envExists := os.LookupEnv("PORT")
	if !envExists {
		port = "8080"
	}

	app.Listen(":" + port)
}
