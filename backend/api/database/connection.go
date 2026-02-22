package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func Connect() {
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()

	mongoURI := "mongodb://localhost:27017"
	Client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		// panic(err)
		fmt.Println("error connect to db", err.Error())
		return
	}
	fmt.Println("Connected to MongoDB")
	DB = Client.Database("socialapp")
}
