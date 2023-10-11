package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func DBSEt() *mongo.Client {
	client, err:= mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	
	if err!= nil {
		log.Fatal(err)
	}
	ctx, canclel :=context.WithTimeout(context.Background(), 100 *time.Second)

	defer canclel()

	err = client.Connect(ctx)
	if err!= nil{
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err!= nil{
		log.Println("failed to connect to mongodb :")
		return nil
	}

	fmt.Println("Successfull connected to database")
	return client
}

var Client *mongo.Client = DBSEt()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection{
	 var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	 return collection
}
func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var produtCollection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)
	return produtCollection
}