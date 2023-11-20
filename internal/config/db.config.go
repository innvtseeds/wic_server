package config

import (
	"context"
	"log"
	"os"
	"time"

	customLogger "github.com/innvtseeds/wdic-server/library/logger"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var myLogger = customLogger.NewLogger()

// holds the configuration details
type MongoDBConfig struct {
	ConnectionUrl string
}

// holds the mongo connection client
type MongoDBConnection struct {
	Client              *mongo.Client
	isConnectionSuccess bool
}

var DbConnection mongo.Database

func NewMongoDBConnection(config *MongoDBConfig) (*MongoDBConnection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(config.ConnectionUrl)

	client, err := mongo.Connect(ctx, clientOptions)

	DbConnection = *client.Database(os.Getenv("DB_NAME"))

	if err != nil {
		return nil, err
	}

	return &MongoDBConnection{
		Client:              client,
		isConnectionSuccess: true,
	}, nil
}

func LoadDBConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	connectionUrl := os.Getenv("DB_URL")

	if connectionUrl == "" {
		myLogger.Error("NO DB URL")
		panic("No DB URL")
	}

	connConfig := &MongoDBConfig{
		ConnectionUrl: connectionUrl,
	}

	connection, err := NewMongoDBConnection(connConfig)

	if err != nil {
		myLogger.Error("Mongo Connection Failed")
	}

	if connection.isConnectionSuccess {
		myLogger.Info("MONGO DB SUCCESSFULLY CONNECTED")
	}
}
