package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/KazakovDenis/contra/internal/common/log"
	"github.com/KazakovDenis/contra/internal/contrad/contants"
	"github.com/KazakovDenis/contra/internal/contrad/database"
)

type ProviderRepo struct {
	Repo
}

type ProviderDoc struct {
	Name string
}

func NewProviderRepo(ctx *context.Context) *ProviderRepo {
	return &ProviderRepo{
		Repo{
			ctx: ctx,
		},
	}
}

func (repo *ProviderRepo) Add(name string) (string, error) {
	db := (*repo.ctx).Value(contants.Database).(*mongo.Database)
	coll := db.Collection(database.CollProviders)
	result, err := coll.InsertOne(*repo.ctx, ProviderDoc{Name: name})
	if err != nil {
		return "ERROR", err
	}
	objectId := result.InsertedID.(primitive.ObjectID).Hex()
	log.Info("New provider has been added: %s[%s]", name, objectId)
	return objectId, nil
}
