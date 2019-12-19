package db

import (
	"context"
	"time"

	"github.com/aglyzov/go-patch"
	"github.com/cottrellio/cottrellio_go/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const userCollection string = "users"

// UserCreate creates a user.
func (m *MongoDB) UserCreate(user model.User) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	collection := m.client.Database(mongoDatabase).Collection(userCollection)

	user.CreatedAt = time.Now()

	// Create user in DB.
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	// Update instance with ID.
	if id, ok := res.InsertedID.(primitive.ObjectID); ok {
		user.ID = id
	}

	return &user, nil
}

// UserList lists users by filters and options.
func (m *MongoDB) UserList(filters map[string][]string, opts map[string]string) ([]*model.User, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	collection := m.client.Database(mongoDatabase).Collection(userCollection)

	// Build query from filters.
	query, err := queryBuilder(filters)
	if err != nil {
		return nil, -1, err
	}

	// Build find options from opts.
	findOptions, err := optionsBuilder(opts)
	if err != nil {
		return nil, -1, err
	}

	// Get total count of items that match query.
	totalItems, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return nil, -1, err
	}

	// Find list of users that match query.
	cursor, err := collection.Find(ctx, query, findOptions)
	if err != nil {
		return nil, -1, err
	}

	// Decode users and append to response list.
	var users []*model.User
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user model.User
		cursor.Decode(&user)
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, -1, err
	}

	return users, totalItems, nil
}

// UserDetail retrieves a user by ID.
func (m *MongoDB) UserDetail(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	collection := m.client.Database(mongoDatabase).Collection(userCollection)

	// Convert id to ObjectID.
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Build filter.
	filter := bson.M{
		"_id": _id,
	}

	// Find user from DB.
	var user model.User
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserUpdate updates partial/full user.
func (m *MongoDB) UserUpdate(id string, changes model.User) (*model.User, error) {
	// Configure context with a set timeout. Cancel if exceeds timeout.
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	collection := m.client.Database(mongoDatabase).Collection(userCollection)

	// Get objectID from id.
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Get user from mongo.
	var user model.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	// Update user with changes.
	_, err = patch.Struct(&user, changes)
	if err != nil {
		return nil, err
	}

	// Create mongo update bson object.
	update := bson.M{
		"$set": bson.M{
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
		},
	}

	// Setup filter for doc being updated.
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	// Save to mongo.
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UserDelete deletes a user.
func (m *MongoDB) UserDelete(id string) error {
	// Configure context with a set timeout. Cancel if exceeds timeout.
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	collection := m.client.Database(mongoDatabase).Collection(userCollection)

	// Get objectID from id.
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Delete from mongo.
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}
