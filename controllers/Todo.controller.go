package controllers

import (
	"github.com/Omsmir/Go-Fiber/models"
	"github.com/Omsmir/Go-Fiber/services"
	"github.com/Omsmir/Go-Fiber/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var StatusDetector = utils.StatusDetector

func GetTodoHandler(c *fiber.Ctx) error {
	todos,err := services.GetTodo()

	if err != nil {
		return StatusDetector(c,fiber.NewError(fiber.StatusInternalServerError,"bad request"))
	}

	return c.Status(200).JSON(todos)
}

func CreateTodo (c *fiber.Ctx) error {

	var todo models.Todo 

	if err := c.BodyParser(&todo); err != nil {
		return StatusDetector(c,fiber.NewError(fiber.StatusBadRequest,"bad request"))
	}


	if todo.Title == "" {
		return StatusDetector(c,fiber.NewError(fiber.StatusNotFound,"some fields are missing"))
	}

	result,err := services.CreateTodo(todo)
	if err != nil {
		return StatusDetector(c,fiber.NewError(fiber.StatusInternalServerError,"error creating a document"))
	}
	return c.Status(201).JSON(result)
}

func GetSingleTodoHandler (c *fiber.Ctx) error {
	var id = c.Params("id")

	todo,err := services.GetSingleTodo(id)

	if err != nil {
		return StatusDetector(c,fiber.NewError(fiber.StatusInternalServerError,err.Error()))
	}

	return c.Status(200).JSON(todo)
}


func DeleteSingleTodo (c *fiber.Ctx) error {
	id := c.Params("id")

	_,err := services.GetSingleTodo(id)
	if err != nil {
		return StatusDetector(c,fiber.NewError(fiber.StatusNotFound,err.Error()))
	}

	result,err := services.DeleteTodo(id)
	if (err != nil) && (result.DeletedCount == 0) {
		return StatusDetector(c,fiber.NewError(fiber.StatusInternalServerError,err.Error()))
	}

	return c.Status(200).JSON(fiber.Map{"message":"todo deleted successfully"})
}

func UpdateTodoHandler (c *fiber.Ctx) error {

	var todo models.Todo
	 id := c.Params("id")

	 if err := c.BodyParser(&todo); err != nil {
		return StatusDetector(c,fiber.NewError(fiber.StatusBadRequest,"bad request"))
	}

	 _,err := services.GetSingleTodo(id)
	 if err != nil {
		return StatusDetector(c,fiber.NewError(fiber.StatusNotFound,err.Error()))
	 }
	 
	 update := bson.M{"$set":bson.M{
		"title":todo.Title,
		"completed":todo.Completed,
	 }}

	 result,err := services.UpdateTodo(id,update)
	 if (err != nil ) && (result.MatchedCount == 0){
		return StatusDetector(c,fiber.NewError(fiber.StatusInternalServerError,"error updating todo"))
	 }

	 return c.Status(200).JSON(fiber.Map{"message":"todo updated successfully","update":update})

}