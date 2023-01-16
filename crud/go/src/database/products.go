package database

import (
	"context"
	"crud/src/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.TODO()

func GetAll() ([]models.Product, error) {

	var products []models.Product

	cur, _ := Collection.Find(ctx, bson.D{})

	defer cur.Close(ctx)

	err := cur.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetById(id string) (models.Product, error) {

	objectId, _ := primitive.ObjectIDFromHex(id)

	var product models.Product

	err := Collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&product)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found")
	}
	if err != nil {
		return product, err
	}

	return product, nil
}

func Create(product models.Product) (string, error) {

	insert, err := Collection.InsertOne(ctx, product)
	if err != nil {
		return "", err
	}

	return insert.InsertedID.(primitive.ObjectID).Hex(), nil
}
