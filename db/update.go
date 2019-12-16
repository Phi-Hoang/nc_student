package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// AddStudent to add a student into DB
func AddStudent(student *Student) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, student)
	return res, err
}

// UpdateStudent to update information of a student
func UpdateStudent(student *StudentUpdateRequest) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"email": student.Email}
	update := bson.M{"$set": student}
	res, err := collection.UpdateOne(ctx, filter, update)
	return res, err
}
