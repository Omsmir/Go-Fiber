package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Omsmir/Go-Fiber/config"
	"github.com/Omsmir/Go-Fiber/middleware"
	"github.com/Omsmir/Go-Fiber/routes"
	"github.com/Omsmir/Go-Fiber/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

var pl = fmt.Println

var collection *mongo.Collection

var router = routes.Router

func main() {

	err := godotenv.Load(".env")

	utils.ErrorCheck("error loading .env file", err)

	port := os.Getenv("PORT")

	app := fiber.New()

	config.MongoConnection()


	router(app)
	app.Use(middleware.NotFound)


	log.Fatal(app.Listen(":" + port))

}

