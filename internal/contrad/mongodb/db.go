package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, uri string, name string) (*mongo.Database, func()) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	disconnect := func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}
	return client.Database(name), disconnect
}
