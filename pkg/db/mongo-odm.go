package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbTimeout = 10 * time.Second
	pageLimit = 10 // TODO: kept it small for testing, change whenever appropriate.
)

var mongoDatabase string

// MongoDB stores mongo client used by repository.
type MongoDB struct {
	client *mongo.Client
}

// Connect Establish connection to the db using environment vars.
func (m *MongoDB) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	uri := buildMongoURI()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	log.Printf("Application connected to MongoDB (%s)...", uri)
	m.client = client

	return nil
}

// Mongo URI pattern - mongodb://[username:password@]host[:port]/[database]
// MONGO_USERNAME - Mongo Username
// MONGO_PASSWORD - Mongo Password
// MONGO_HOST - Mongo Host
// MONGO_PORT - Mongo Port
// MONGO_DATABASE - Mongo Database
func buildMongoURI() string {
	// Determine protocol
	mongoProtocol := "mongodb+srv"
	if len(os.Getenv("MONGO_PROTOCOL")) > 0 {
		mongoProtocol = os.Getenv("MONGO_PROTOCOL")
	}

	// Determine mongo Auth.
	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoAuth := ""
	if len(mongoUsername) > 0 && len(mongoPassword) > 0 {
		mongoAuth = fmt.Sprintf("%s:%s", mongoUsername, mongoPassword)
	}

	// Determine mongo Host.
	mongoHost := os.Getenv("MONGO_HOST")

	// Determine mongo Port.
	mongoPort := ""
	if len(os.Getenv("MONGO_PORT")) > 0 {
		mongoPort = fmt.Sprintf(":%s", os.Getenv("MONGO_PORT"))
	}

	// Determine mongo Database.
	mongoDatabase = os.Getenv("MONGO_DATABASE")

	// Determine Mongo querystring.
	mongoQuerystring := ""
	if len(os.Getenv("MONGO_QUERYSTRING")) > 0 {
		mongoQuerystring = fmt.Sprintf("?%s", os.Getenv("MONGO_QUERYSTRING"))
	}

	// Build mongo URI.
	return fmt.Sprintf(`%s://%s@%s%s/%s%s`, mongoProtocol, mongoAuth, mongoHost, mongoPort, mongoDatabase, mongoQuerystring)
}
