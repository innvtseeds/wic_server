package repository

import "go.mongodb.org/mongo-driver/mongo"

func CollectionManager() *mongo.Collection {
	return GetCollection("ingredient")
}

func CreateIngredient() {
}

func UpdateIngredient() {
}

func DeleteIngredient() {
}

func GetIngredient() {
}

func GetIngredients() {
}
