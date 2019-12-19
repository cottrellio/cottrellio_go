package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/cottrellio/cottrellio_go/pkg/creating"
	"github.com/cottrellio/cottrellio_go/pkg/deleting"
	"github.com/cottrellio/cottrellio_go/pkg/model"
	"github.com/cottrellio/cottrellio_go/pkg/reading"
	"github.com/cottrellio/cottrellio_go/pkg/updating"
	"github.com/gorilla/mux"
)

// PostCreateEndpoint creates a post.
func PostCreateEndpoint(creator creating.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var post model.Post
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			log.Printf("Error decoding request body: %v: %s", r, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		created, err := creator.PostCreate(post)
		if err != nil {
			log.Println("Error creating post:", post, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	}
}

// PostListEndpoint gets list of posts.
func PostListEndpoint(reader reading.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := r.URL.Query()
		items, totalItems, err := reader.PostList(params)
		if err != nil {
			log.Printf("Error listing posts: %v: %s", params, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := map[string]interface{}{"items": items, "total_items": totalItems}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

// PostDetailEndpoint gets post detail.
func PostDetailEndpoint(reader reading.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := mux.Vars(r)["id"]
		post, err := reader.PostDetail(id)
		if err != nil {
			log.Printf("Error getting post detail: %s: %s", id, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	}
}

// PostUpdateEndpoint updates a post.
func PostUpdateEndpoint(updater updating.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := mux.Vars(r)["id"]

		var post model.Post
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			log.Printf("Error decoding request: %v: %s", r, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Patch post.
		updated, err := updater.PostUpdate(id, post)
		if err != nil {
			log.Println("Error updating a user:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updated)
	}
}

// PostDeleteEndpoint deletes a post.
func PostDeleteEndpoint(deleter deleting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := mux.Vars(r)["id"]

		// Delete a post.
		err := deleter.PostDelete(id)
		if err != nil {
			log.Println("Error deleting a post:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
