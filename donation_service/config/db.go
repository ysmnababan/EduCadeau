package config

import (
	"context"
	"fmt"

	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, dbname string) (*mongo.Client, *mongo.Database) {
	// get uri from .env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	mongoURI := os.Getenv("MONGO_URI")
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	db := client.Database(dbname)

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(os.Stdout, []any{"Pinged your deployment. You successfully connected to MongoDB!"}...)

	return client, db
}
