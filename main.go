package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var pl = fmt.Println

type TODO struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Todo      string `json:"todo"`
}

func main() {

	app := fiber.New()

	Todos := []TODO{}


	log.Fatal(app.Listen("localhost:8080"))
}
