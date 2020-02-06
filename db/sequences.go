package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetNextID(col *mongo.Collection, sequenceName string) (int, error) {
	filter := bson.M{"_id": sequenceName}
	update := bson.M{"$inc": bson.M{"sequence_value": 1}}
	option := options.FindOneAndUpdate()
	option.SetUpsert(true)
	var sequence Sequence
	err := col.FindOneAndUpdate(context.TODO(), filter, update, option).Decode(&sequence)

	return sequence.Value, err
}
