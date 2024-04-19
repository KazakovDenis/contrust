package repo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/KazakovDenis/contra/internal/contrad/constants"
	"github.com/KazakovDenis/contra/internal/contrad/local_errors"
	"github.com/KazakovDenis/contra/internal/contrad/mongodb"
)

type MongoDbRepo struct {
	ctx            *context.Context
	collectionName string
}

type Document interface {
}

func (repo *MongoDbRepo) getCollection() *mongo.Collection {
	db := (*repo.ctx).Value(constants.Database).(*mongo.Database)
	return db.Collection(repo.collectionName)
}

func (repo *MongoDbRepo) insert(doc Document) (string, error) {
	coll := repo.getCollection()
	result, err := coll.InsertOne(*repo.ctx, doc)
	if err == nil {
		return mongodb.GetObjectId(result), nil
	}

	var writeException mongo.WriteException

	if errors.As(err, &writeException) {
		return "WRITE_ERROR", &local_errors.DatabaseWriteError{}
	}
	return "ERROR", err
}
