package v1

import (
	"net/http"

	"github.com/cottrellio/cottrellio_go/pkg/creating"
	"github.com/cottrellio/cottrellio_go/pkg/deleting"
	"github.com/cottrellio/cottrellio_go/pkg/reading"
	"github.com/cottrellio/cottrellio_go/pkg/updating"
	"github.com/gorilla/mux"
)

// Handler handles all API V1 routes.
func Handler(router *mux.Router, c creating.Service, r reading.Service, u updating.Service, d deleting.Service) **mux.Router {
	v1 := router.PathPrefix("/v1").Subrouter()

	// Handle V1 routes
	v1.HandleFunc("", IndexRoute())
	v1.HandleFunc("/", IndexRoute())
	// Users.
	v1.HandleFunc("/users", UserCreateEndpoint(c)).Methods("POST")
	v1.HandleFunc("/users/", UserCreateEndpoint(c)).Methods("POST")
	v1.HandleFunc("/users", UserListEndpoint(r)).Methods("GET")
	v1.HandleFunc("/users/", UserListEndpoint(r)).Methods("GET")
	v1.HandleFunc("/users/{id}", UserDetailEndpoint(r)).Methods("GET")
	v1.HandleFunc("/users/{id}/", UserDetailEndpoint(r)).Methods("GET")
	v1.HandleFunc("/users/{id}", UserUpdateEndpoint(u)).Methods("PATCH")
	v1.HandleFunc("/users/{id}/", UserUpdateEndpoint(u)).Methods("PATCH")
	v1.HandleFunc("/users/{id}", UserDeleteEndpoint(d)).Methods("DELETE")
	v1.HandleFunc("/users/{id}/", UserDeleteEndpoint(d)).Methods("DELETE")
	// Posts.
	v1.HandleFunc("/posts", PostCreateEndpoint(c)).Methods("POST")
	v1.HandleFunc("/posts/", PostCreateEndpoint(c)).Methods("POST")
	v1.HandleFunc("/posts", PostListEndpoint(r)).Methods("GET")
	v1.HandleFunc("/posts/", PostListEndpoint(r)).Methods("GET")
	v1.HandleFunc("/posts/{id}", PostDetailEndpoint(r)).Methods("GET")
	v1.HandleFunc("/posts/{id}/", PostDetailEndpoint(r)).Methods("GET")
	v1.HandleFunc("/posts/{id}", PostUpdateEndpoint(u)).Methods("PATCH")
	v1.HandleFunc("/posts/{id}/", PostUpdateEndpoint(u)).Methods("PATCH")
	v1.HandleFunc("/posts/{id}", PostDeleteEndpoint(d)).Methods("DELETE")
	v1.HandleFunc("/posts/{id}/", PostDeleteEndpoint(d)).Methods("DELETE")

	return &v1
}

// IndexRoute handles all index API requests.
func IndexRoute() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API / V1 :)"))
	}
}
