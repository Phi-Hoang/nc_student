package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Phi-Hoang/nc_student/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func Test() {
	fmt.Println("connect & insert db")
}
func init() {
	connect()
	insertNunber()
}

func connect() {
	fmt.Println(config.Config.Mongo.URI)
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.Mongo.URI))
	if err != nil {
		log.Fatalf("create client error: %v", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Printf("ping error: %v\n", err)
	}

	Client = client
}

func insertNunber() {
	collection := Client.Database("testing").Collection("numbers")
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		log.Fatalf("test inserting number error: %v", err)
	}
	id := res.InsertedID
	fmt.Println(id)
}
