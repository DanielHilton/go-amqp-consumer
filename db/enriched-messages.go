package db

import (
	"context"
	"fmt"
	"github.com/DanielHilton/go-amqp-consumer/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetSample() ([]structs.EnrichedMessage, error) {
	coll := MongoClient.Database("poc").Collection("go")
	docs, err := coll.Aggregate(context.Background(), mongo.Pipeline{
		bson.D{
			{"$sample", bson.D{
				{"size", 3},
			}},
		}}, nil)

	var results []structs.EnrichedMessage
	if err != nil {
		_ = fmt.Errorf("failed to aggregate: %w", err)
		return results, err
	}

	err = docs.All(context.Background(), &results)
	return results, err
}

func InsertEnrichedMessage(message structs.EnrichedMessage) error {
	coll := MongoClient.Database("poc").Collection("go")
	_, err := coll.InsertOne(context.Background(), message)

	return err
}
