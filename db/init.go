package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/phihdn/nc_student/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client common DB Client
var Client *mongo.Client

const (
	//DbName is Database Name
	DbName = "go-learning"
	//ColName is Collection Name
	ColName = "students"
)

// Test to test DB by insterting Pi number into test db
func Test() {
	insertNunber()
	fmt.Println("connect & insert db")
}
func init() {
	connect()
	//insertNunber()
}

func connect() {
	fmt.Println(config.Config.Mongo.URI)

	clientOptions := options.Client().ApplyURI(config.Config.Mongo.URI)
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(4)
	clientOptions.SetReadPreference(readpref.Nearest())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("connect error: %v", err)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Printf("ping error: %v\n", err)
	}

	Client = client
}

func insertNunber() {
	collection := Client.Database("testing").Collection("numbers")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		log.Fatalf("test inserting number error: %v", err)
	}
	id := res.InsertedID
	fmt.Println(id)
}
