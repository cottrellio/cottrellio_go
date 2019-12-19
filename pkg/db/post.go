package db

import (
	"context"
	"time"

	"github.com/aglyzov/go-patch"
	"github.com/cottrellio/cottrellio_go/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const postCollection string = "posts"

// PostCreate creates a post.
func (m *MongoDB) PostCreate(post model.Post) (*model.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	collection := m.client.Database(mongoDatabase).Collection(postCollection)

	post.CreatedAt = time.Now()

	// Create post in DB.
	res, err := collection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}

	// Update instance with ID.
	if id, ok := res.InsertedID.(primitive.ObjectID); ok {
		post.ID = id
	}

	return &post, nil
}

// PostList lists posts by filters and options.
func (m *MongoDB) PostList(filters map[string][]string, opts map[string]string) ([]*model.Post, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	collection := m.client.Database(mongoDatabase).Collection(postCollection)

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

	// Decode posts and append to response list.
	var posts []*model.Post
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post model.Post
		cursor.Decode(&post)
		posts = append(posts, &post)
	}
	if err := cursor.Err(); err != nil {
		return nil, -1, err
	}

	return posts, totalItems, nil
}

// PostDetail retrieves a post by ID.
func (m *MongoDB) PostDetail(id string) (*model.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	collection := m.client.Database(mongoDatabase).Collection(postCollection)

	// Convert id to ObjectID.
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Build filter.
	filter := bson.M{
		"_id": objID,
	}

	// Find post from DB.
	var post model.Post
	err = collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// PostUpdate updates partial/full post.
func (m *MongoDB) PostUpdate(id string, changes model.Post) (*model.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	collection := m.client.Database(mongoDatabase).Collection(postCollection)

	// Get objectID from id.
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	// Get post from mongo.
	var post model.Post
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		return nil, err
	}

	// Update post with changes.
	_, err = patch.Struct(&post, changes)
	if err != nil {
		return nil, err
	}

	// Create mongo update bson object.
	update := bson.M{
		"$set": bson.M{
			"title":   post.Title,
			"slug":    post.Slug,
			"body":    post.Body,
			"excerpt": post.Excerpt,
			"author":  post.Author,
			"status":  post.Status,
		},
	}

	// Setup filter for doc being updated.
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	// Save to mongo.
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// PostDelete deletes a post.
func (m *MongoDB) PostDelete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	collection := m.client.Database(mongoDatabase).Collection(postCollection)

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
