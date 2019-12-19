package api

import (
	"net/http"

	"github.com/cottrellio/cottrellio_go/pkg/creating"
	"github.com/cottrellio/cottrellio_go/pkg/deleting"
	"github.com/cottrellio/cottrellio_go/pkg/reading"
	v1 "github.com/cottrellio/cottrellio_go/pkg/routing/api/v1"
	"github.com/cottrellio/cottrellio_go/pkg/updating"
	"github.com/gorilla/mux"
)

// Handler handles all API routes.
func Handler(router *mux.Router, c creating.Service, r reading.Service, u updating.Service, d deleting.Service) **mux.Router {
	api := router.PathPrefix("/api").Subrouter()

	// Handle API routes
	api.HandleFunc("", IndexRoute())
	api.HandleFunc("/", IndexRoute())
	// api.Use(authing.JWTMiddleware)
	v1.Handler(api, c, r, u, d)

	return &api
}

// IndexRoute handles all index API requests.
func IndexRoute() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API :)"))
	}
}
