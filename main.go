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

	app.Get("/api/todos", func(c *fiber.Ctx) error {

		return c.Status(200).JSON(Todos)
	})

	app.Post("/api/todo", func(c *fiber.Ctx) error {

		todo := &TODO{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Todo == "" {
			return c.Status(404).JSON(fiber.Map{"message": "some fields are missing"})
		}

		todo.ID = len(Todos) + 1

		Todos = append(Todos, *todo)

		return c.Status(201).JSON(todo)
	})

	app.Put("/api/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, Todo := range Todos {
			if fmt.Sprint(Todos[i].ID) == id {
				Todos[i].Completed = true
				return c.Status(200).JSON(Todo)

			}
		}
		return c.Status(400).JSON(fiber.Map{"message": "err encountred"})
	})
	log.Fatal(app.Listen("localhost:8080"))
}
