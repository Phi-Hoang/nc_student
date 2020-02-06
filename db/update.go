package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// AddStudent to add a student into DB
func AddStudent(req *StudentAddRequest) (interface{}, error) {
	sequenceCol := Client.Database(DbName).Collection("sequences")
	id, err := GetNextID(sequenceCol, "studentId")
	if err != nil {
		return nil, err
	}
	student := Student{
		ID:        id,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Age:       req.Age,
		ClassName: req.ClassName,
		Email:     req.Email,
	}

	collection := Client.Database(DbName).Collection(ColName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, student)
	if err != nil {
		return nil, err
	}

	return res, nil
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
