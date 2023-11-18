package repository

import (
	"github.com/innvtseeds/wdic-server/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(collectionName string) *mongo.Collection {
	return config.DbConnection.Collection(collectionName)
}
