package mongo

import (
	"context"
	"errors"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/geeksheik9/common-functions/pkg/models"
)

// Mongo client struct that holds cleint and database names
type MongoDB struct {
	Client        *mongo.Client
	DatabaseNames map[string]string
}

// Initializing a mongo client using contexts and configs
func InitializeClient(ctx context.Context, uri string, config *models.Config) (*mongo.Client, error) {
	options := options.Client().ApplyURI(uri)

	if strings.Trim(options.GetURI(), " ") == "" {
		return nil, errors.New("MONGO_URI is not set")
	}

	err := options.Validate()
	if err != nil {
		return nil, err
	}

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, err
}

// InitializeDatabases Factory for the dao implementation. Returns a dao connected to the designated MongoDB database for DB operations.
// The database connection is made using configuration in the config.go file
func (db *MongoDB) InitializeDatabases(client *mongo.Client, config *models.Config) *MongoDB {
	return &MongoDB{
		Client:        client,
		DatabaseNames: config.Databases,
	}
}
