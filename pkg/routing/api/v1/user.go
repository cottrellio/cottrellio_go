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

// UserCreateEndpoint creates a user.
func UserCreateEndpoint(creator creating.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("Error decoding request body: %v: %s", r, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		created, err := creator.UserCreate(user)
		if err != nil {
			log.Println("Error creating user:", user, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	}
}

// UserListEndpoint gets list of users.
func UserListEndpoint(reader reading.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := r.URL.Query()
		items, totalItems, err := reader.UserList(params)
		if err != nil {
			log.Printf("Error listing users: %v: %s", params, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res := map[string]interface{}{"items": items, "total_items": totalItems}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

// UserDetailEndpoint gets user detail.
func UserDetailEndpoint(reader reading.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := mux.Vars(r)["id"]
		user, err := reader.UserDetail(id)
		if err != nil {
			log.Printf("Error getting user detail: %s: %s", id, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

// UserUpdateEndpoint updates a user.
func UserUpdateEndpoint(updater updating.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := mux.Vars(r)["id"]

		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("Error decoding request: %v: %s", r, err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Patch user.
		updated, err := updater.UserUpdate(id, user)
		if err != nil {
			log.Println("Error updating a user:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updated)
	}
}

// UserDeleteEndpoint deletes a user.
func UserDeleteEndpoint(deleter deleting.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := mux.Vars(r)["id"]

		// Delete a user.
		err := deleter.UserDelete(id)
		if err != nil {
			log.Println("Error deleting a user:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
