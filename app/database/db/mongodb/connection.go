package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

//DBInstance func
func DBInstance(configMongo string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(configMongo))
	if err != nil {
		panic("Failed to create a connection to database mongodb" + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic("Failed to create a connection to database mongodb" + err.Error())
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

//OpenCollection is a  function makes a connection with a collection in the database
func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	Database := os.Getenv("MONGODB_DATABASE")
	var collection *mongo.Collection = client.Database(Database).Collection(collectionName)
	return collection
}
