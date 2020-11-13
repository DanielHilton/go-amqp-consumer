package db

import (
	"context"
	"fmt"
	H "github.com/DanielHilton/go-amqp-consumer/helpers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"time"
)

var MongoClient *mongo.Client

func Init() (context.Context, context.CancelFunc) {
	mc, err := mongo.NewClient(options.Client().ApplyURI(H.EnvVarOrFallback("MONGO_URI", "mongodb://localhost:27017")))
	if err != nil {
		log.Fatal(err)
	}

	MongoClient = mc

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = MongoClient.Connect(ctx)
	H.ExitOnFail(err, "failed to connect to MongoDB")

	fmt.Println("Connected to MongoDB")
	return ctx, cancel
}
