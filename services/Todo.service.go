package services

import (
	"context"
	"errors"
	"log"

	"github.com/Omsmir/Go-Fiber/config"
	"github.com/Omsmir/Go-Fiber/models"
	"github.com/Omsmir/Go-Fiber/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


var errorCheck = utils.ErrorCheck


func GetSingleTodo(id string) (*models.Todo,error){

	var collection = config.GetCollection("todos")

	objId,err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil,errors.New("Invalid ObjectId") 
	}
	var existingTodo models.Todo
	err = collection.FindOne(context.Background(),bson.M{"_id":objId}).Decode(&existingTodo)

	if err != nil {
		return nil,errors.New("Todo Is Not Found")
	}

	return &existingTodo,nil

}

func GetTodo() ([]models.Todo,error) {

	var todos []models.Todo

	var collection = config.GetCollection("todos")

	cursor,err := collection.Find(context.Background(),bson.M{})

	if err != nil {
		return nil,err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var todo models.Todo
		if err  := cursor.Decode(&todo); err != nil {
			log.Println("Error decoding todo:", err)

		}else {
			todos = append(todos, todo)
		}
	}

	return todos,nil
}


func CreateTodo(todo models.Todo)(*mongo.InsertOneResult,error){
	var collection = config.GetCollection("todos")

	return collection.InsertOne(context.Background(),todo) 
}

func DeleteTodo(id string)(*mongo.DeleteResult,error){
	collection := config.GetCollection("todos")
	objId,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil,err
	}
	return collection.DeleteOne(context.Background(),bson.M{"_id":objId})
}


func UpdateTodo(id string,update any) (*mongo.UpdateResult,error){
	collection := config.GetCollection("todos")

	objId,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil,err
	}
	
	return collection.UpdateOne(context.Background(),bson.M{"_id":objId},update)
}