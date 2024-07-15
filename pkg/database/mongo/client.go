package mongo

import (
	"context"
	"errors"
	"net/url"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/geeksheik9/common-functions/pkg/models"
)

// Mongo client struct that holds cleint and database names
type MongoDB struct {
	Client    *mongo.Client
	Databases map[string]string
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
func (db *MongoDB) InitializeDatabases(client *mongo.Client, databases []string) *MongoDB {
	dbMap := make(map[string]string)
	for _, database := range databases {
		dbMap[database] = database
	}
	return &MongoDB{
		Client:    client,
		Databases: dbMap,
	}
}

// BuildQuery sets up the mongo query
func BuildQuery(ID *primitive.ObjectID, name *string, params ...bson.M) bson.M {
	conditions := []bson.M{}
	c := bson.M{}

	if ID != nil {
		conditions = append(conditions, bson.M{"_id": ID})
	}

	if name != nil {
		conditions = append(conditions, bson.M{"name": &name})
	}

	if len(params) > 0 {
		for _, otherFilter := range params {
			conditions = append(conditions, otherFilter)
		}
	}

	if len(conditions) != 0 {
		c = bson.M{"$and": conditions}
	}
	return c
}

// BuildFilter sets up the mongo filtering
func BuildFilter(queryParams url.Values) (int, int, string, bson.M) {
	filter := bson.M{}
	filters := []bson.M{}
	pageNumber := 0
	pageCount := 10000
	sort := "priority"
	if len(queryParams) > 0 {
		for queryParam, paramValue := range queryParams {
			switch queryParam {
			case "pageNumber":
				pageNumber, _ = strconv.Atoi(paramValue[0])
			case "pageCount":
				pageCount, _ = strconv.Atoi(paramValue[0])
			case "sort":
				sort = paramValue[0]
			default:
				m := bson.M{queryParam: paramValue[0]}
				filters = append(filters, m)
			}
		}
		if len(filters) > 0 {
			filter = bson.M{"$and": filters}
		} else {
			filter = nil
		}
		return pageNumber, pageCount, sort, filter
	}

	return pageNumber, pageCount, sort, nil
}
