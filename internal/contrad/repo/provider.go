package repo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/KazakovDenis/contra/internal/contrad/contants"
	"github.com/KazakovDenis/contra/internal/contrad/database"
)

type ProviderRepo struct {
	Repo
}

func NewProviderRepo(ctx *context.Context) *ProviderRepo {
	return &ProviderRepo{
		Repo{
			ctx: ctx,
		},
	}
}

func (repo *ProviderRepo) Add(name string) error {
	db := (*repo.ctx).Value(contants.Database).(*mongo.Database)
	coll := db.Collection(database.CollProviders)
	result, err := coll.InsertOne(*repo.ctx, struct{ name string }{name: name})
	if err != nil {
		return err
	}
	log.Printf("Added new provider: %s", result.InsertedID)
	return nil
}
