package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const defaultLimit int64 = 10

// QueryBuilder builds a mongo query from filters.
func queryBuilder(filters map[string][]string) (bson.M, error) {
	return bson.M{}, nil
}

// optionsBuilder builds mongo findOptions from opts.
func optionsBuilder(opts map[string]string) (*options.FindOptions, error) {
	findOptions := options.Find()
	findOptions.SetLimit(defaultLimit)

	return findOptions, nil
}
