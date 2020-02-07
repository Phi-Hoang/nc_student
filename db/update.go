package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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

func DeleteStudent(req *StudentDeleteRequest) (interface{}, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var filter bson.M
	bs, err := json.Marshal(req)
	err = json.Unmarshal(bs, &filter)
	if err != nil {
		log.Printf("marshal error: %v", err)
	}

	fmt.Println("Filter", filter)

	noDeletedStudents, err := Client.Database(DbName).Collection(ColName).DeleteMany(ctx, filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	} else {
		log.Println(noDeletedStudents.DeletedCount, " students have been deleted.")
	}
	return noDeletedStudents.DeletedCount, nil
}

func DeleteStudentById(id int) (*Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": id}

	var student Student
	err := Client.Database(DbName).Collection(ColName).FindOneAndDelete(ctx, filter).Decode(&student)
	if err != nil {
		log.Printf("delete error: %v", err)
		return nil, err
	}

	return &student, nil
}