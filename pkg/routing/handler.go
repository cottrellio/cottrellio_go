package routing

import (
	"net/http"

	"github.com/cottrellio/cottrellio_go/pkg/creating"
	"github.com/cottrellio/cottrellio_go/pkg/deleting"
	"github.com/cottrellio/cottrellio_go/pkg/reading"
	"github.com/cottrellio/cottrellio_go/pkg/routing/api"
	"github.com/cottrellio/cottrellio_go/pkg/updating"
	"github.com/gorilla/mux"
)

// Handler handles routes.
func Handler(c creating.Service, r reading.Service, u updating.Service, d deleting.Service) http.Handler {
	router := mux.NewRouter()

	// Handle root routes.
	router.HandleFunc("", IndexRoute())
	router.HandleFunc("/", IndexRoute())

	// Handle API routes.
	api.Handler(router, c, r, u, d)

	return router
}

// IndexRoute handles all root requests.
func IndexRoute() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(":)"))
	}
}
