package routes

import (
	"github.com/Omsmir/Go-Fiber/controllers"
	"github.com/gofiber/fiber/v2"
)


func Router(app *fiber.App) {

	router := app.Group("/api")
	router.Get("/todos", controllers.GetTodoHandler)

	router.Get("/todo/:id",controllers.GetSingleTodoHandler)
	router.Post("/todos",controllers.CreateTodo)
	router.Delete("/todo/:id",controllers.DeleteSingleTodo)
	router.Put("/todo/:id",controllers.UpdateTodoHandler)
}
