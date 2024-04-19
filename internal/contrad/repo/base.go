package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/KazakovDenis/contra/internal/contrad/contants"
)

type MongoDbRepo struct {
	ctx            *context.Context
	collectionName string
}

type Document interface {
}

func (repo *MongoDbRepo) getCollection() *mongo.Collection {
	db := (*repo.ctx).Value(contants.Database).(*mongo.Database)
	return db.Collection(repo.collectionName)
}

func (repo *MongoDbRepo) insert(doc Document) (string, error) {
	coll := repo.getCollection()
	result, err := coll.InsertOne(*repo.ctx, doc)
	if err != nil {
		return "ERROR", err
	}
	objectId := result.InsertedID.(primitive.ObjectID).Hex()
	return objectId, nil
}
