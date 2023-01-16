package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Uri            string
	DatabaseName   string
	CollectionName string

	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func (m *MongoDB) Connect() error {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.Uri))

	if err != nil {

		return err
	}

	m.Client = client

	m.Database = m.Client.Database(m.DatabaseName)

	m.Collection = m.Database.Collection(m.CollectionName)

	return nil
}

func (m *MongoDB) Disconnect() {

	m.Client.Disconnect(context.TODO())
}
