package repo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/KazakovDenis/contrust/internal/common/log"
	"github.com/KazakovDenis/contrust/internal/server/constants"
	"github.com/KazakovDenis/contrust/internal/server/local_errors"
	"github.com/KazakovDenis/contrust/internal/server/mongodb"
)

type MongoDbRepo struct {
	ctx            *context.Context
	collectionName string
}

type Document map[string]interface{}

func (repo *MongoDbRepo) getCollection() *mongo.Collection {
	db := (*repo.ctx).Value(constants.Database).(*mongo.Database)
	return db.Collection(repo.collectionName)
}

func (repo *MongoDbRepo) insert(doc interface{}) (string, error) {
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

func (repo *MongoDbRepo) get(provider string, transport *string) ([]Document, error) {
	coll := repo.getCollection()
	filter := Document{
		"provider": provider,
	}
	if transport != nil {
		filter["transport"] = transport
	}

	cursor, err := coll.Find(*repo.ctx, filter)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := cursor.Close(*repo.ctx); err != nil {
			log.Error(err.Error())
		}
	}()

	docs := make([]Document, 0)

	for cursor.Next(*repo.ctx) {
		var doc Document
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	if err := cursor.Err(); err != nil {
		log.Error(err.Error())
	}

	return docs, nil
}
