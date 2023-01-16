package repository

import (
	"compression/src/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompressionRepository struct {
	Collection *mongo.Collection
}

// TODO: search how works that and if i inject or what?
var ctx = context.TODO()

func (d *CompressionRepository) GetAll() ([]models.CompressionDb, error) {

	var compression []models.CompressionDb

	cur, err := d.Collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	err = cur.All(ctx, &compression)

	if err != nil {
		return nil, err
	}

	return compression, nil
}

func (d *CompressionRepository) GetById(id string) (models.CompressionDb, error) {

	var compression models.CompressionDb

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		return compression, err
	}

	err = d.Collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&compression)

	if err == mongo.ErrNoDocuments {

		return compression, nil
	}

	if err != nil {

		return compression, err
	}

	return compression, nil
}

func (d *CompressionRepository) Create(CompressionDb models.CompressionDb) (string, error) {

	insert, err := d.Collection.InsertOne(ctx, CompressionDb)

	if err != nil {

		return "", err
	}

	result := insert.InsertedID.(primitive.ObjectID).Hex()

	return result, nil
}
