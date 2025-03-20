package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Omsmir/Go-Fiber/utils"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB * mongo.Database

func MongoConnection() {
	err := godotenv.Load(".env")
	errorCheck:= utils.ErrorCheck

	errorCheck("error loading .env", err)

	DATABASE := os.Getenv("DATABASE")
	MONGO_URI := os.Getenv("MONGO_URI")
	MONGO_PASSWD := os.Getenv("MONGO_PASSWD")
	MONGO_USER := os.Getenv("MONGO_USER")

	mongoUri := fmt.Sprintf("%s//%s:%s@localhost:27017/%s?authSource=admin", MONGO_URI, MONGO_USER, MONGO_PASSWD, DATABASE)

	CLientOptions := options.Client().ApplyURI(mongoUri)

	client, err := mongo.Connect(context.TODO(), CLientOptions)

	errorCheck("error connecting to mongo", err)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	errorCheck("error connection of mongo", err)

	DB = client.Database(DATABASE)
	fmt.Println("Connected to MongoDB")

}

func GetCollection (collectionName string) (*mongo.Collection){
	return DB.Collection(collectionName)
}