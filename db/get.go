package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// GetAllStudent return all students in DB
func GetAllStudent() (*[]Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{} //map[string]interface{}
	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var students []Student
	err = cur.All(ctx, &students)
	if err != nil {
		log.Printf("cur all error: %v", err)
		return nil, err
	}

	return &students, nil
}
