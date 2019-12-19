package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post model.
type Post struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Slug        string             `json:"slug,omitempty" bson:"slug,omitempty"`
	Body        string             `json:"body,omitempty" bson:"body,omitempty"`
	Excerpt     string             `json:"excerpt,omitempty" bson:"excerpt,omitempty"`
	Author      string             `json:"author,omitempty" bson:"author,omitempty"`
	Status      string             `json:"status,omitempty" bson:"status,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	PublishedAt time.Time          `json:"published_at,omitempty" bson:"published_at,omitempty"`
}
