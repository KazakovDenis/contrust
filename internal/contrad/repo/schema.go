package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/KazakovDenis/contra/internal/common/log"
	"github.com/KazakovDenis/contra/internal/contrad/contants"
)

type SchemaRepo struct {
	Repo
}

type SchemaDoc struct {
	Contract map[string]interface{} `json:"contract"`
}

func NewSchemaRepo(ctx *context.Context) *SchemaRepo {
	return &SchemaRepo{
		Repo{
			ctx: ctx,
		},
	}
}

func (repo *SchemaRepo) Add(collection string, document map[string]interface{}) (string, error) {
	db := (*repo.ctx).Value(contants.Database).(*mongo.Database)
	coll := db.Collection(collection)
	result, err := coll.InsertOne(*repo.ctx, SchemaDoc{Contract: document})
	if err != nil {
		return "ERROR", err
	}
	objectId := result.InsertedID.(primitive.ObjectID).Hex()
	log.Info("New contract has been added: %s[%s]", collection, objectId)
	return objectId, nil
}
