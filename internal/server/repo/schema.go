package repo

import (
	"context"
	"github.com/KazakovDenis/contrust/internal/common/log"
	"github.com/KazakovDenis/contrust/internal/server/mongodb"
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

func (repo *SchemaRepo) Get(provider string, transport *string) ([]Document, error) {
	documents, err := repo.get(provider, transport)
	if err != nil {
		return nil, err
	}
	log.Debug("Contracts found for %s: %s", provider, *transport)
	return documents, nil
}
