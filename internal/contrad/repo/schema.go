package repo

import (
	"context"

	"github.com/KazakovDenis/contra/internal/common/log"
	"github.com/KazakovDenis/contra/internal/contrad/mongodb"
)

type SchemaRepo struct {
	MongoDbRepo
}

type SchemaDoc struct {
	Contract map[string]interface{} `json:"contract"`
}

func NewSchemaRepo(ctx *context.Context) *SchemaRepo {
	return &SchemaRepo{
		MongoDbRepo{
			ctx:            ctx,
			collectionName: mongodb.CollSchemas,
		},
	}
}

func (repo *SchemaRepo) Add(collection string, document map[string]interface{}) (string, error) {
	objectId, err := repo.insert(SchemaDoc{Contract: document})
	if err != nil {
		return "ERROR", err
	}
	log.Info("New contract has been added: %s[%s]", collection, objectId)
	return objectId, nil
}
