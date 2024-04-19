package repo

import (
	"context"

	"github.com/KazakovDenis/contra/internal/common/log"
	"github.com/KazakovDenis/contra/internal/contrad/mongodb"
)

type ProviderRepo struct {
	MongoDbRepo
}

type ProviderDoc struct {
	Name string `json:"name"`
}

func NewProviderRepo(ctx *context.Context) *ProviderRepo {
	return &ProviderRepo{
		MongoDbRepo{
			ctx:            ctx,
			collectionName: mongodb.CollProviders,
		},
	}
}

func (repo *ProviderRepo) Add(name string) (string, error) {
	objectId, err := repo.insert(ProviderDoc{Name: name})
	if err != nil {
		return "ERROR", err
	}
	log.Info("New provider has been added: %s[%s]", name, objectId)
	return objectId, nil
}
