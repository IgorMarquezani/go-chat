package main

import (
	"log"

	"app/adapters/database"
	"app/cmd/routes"
	"app/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("could not load .env:", err)
	}
}

func SetupDatabase() {
	database.Setup()

	db, err := database.GetDBConn(5)
	if err != nil {
		panic(err)
	}

	if err := db.Migrator().AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}
}

func main() {
	SetupEnv()
	SetupDatabase()

	app := fiber.New(fiber.Config{
		BodyLimit:         20 * 1024 * 1024,
		EnablePrintRoutes: true,
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	routes.Setup(app)

	log.Fatal(app.Listen("0.0.0.0:3000"))
}
