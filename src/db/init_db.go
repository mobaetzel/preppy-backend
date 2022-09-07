package db

import (
	"context"
	"fmt"
	"time"

	"github.com/aivot-digital/preppy-backend/src/cfg"
	"github.com/aivot-digital/preppy-backend/src/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Users ICollection[models.User]
var Recipes ICollection[models.Recipe]

func InitDb() func() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	database := client.Database(cfg.DB_DATABASE)

	Users = NewMongoCollection[models.User](database, "users")
	Recipes = NewMongoCollection[models.Recipe](database, "recipes")

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		client.Disconnect(ctx)
	}
}
