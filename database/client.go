package database

import (
	"context"
	"log"
	"otavio-alves/OtakuList/configs"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client .. DB Client
var Client *mongo.Client

// Db ..
var Db *mongo.Database

// CreateClient .. Creates a new client
func CreateClient() {

	// Creates a new client to connect with DB
	Client, err := mongo.NewClient(clientOptions())

	// Checks if any error occurs during the creation of the client
	if err != nil {
		log.Println(configs.FATAL_CLIENT)
		panic(err)
	}

	// Creates a context for client connection function
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	// Calls cancel context function
	defer cancel()

	// Connects with database and checks if any error occurs during
	if err = Client.Connect(ctx); err != nil {
		log.Println(configs.FATAL_CONN_DB)
		panic(err)
	}

	// Gets the wanted database
	Db = Client.Database(configs.DB_NAME)

	log.Print(configs.SUCCESS_CONNECTING_DB)

	return
}

// clientOptions .. Returns the configured client options
func clientOptions() (opts *options.ClientOptions) {

	opts = options.Client()
	opts.ApplyURI(configs.MONGO_HOST)

	return
}
