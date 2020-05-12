package api

import "github.com/gorilla/mux"

// initiate api routes
func (a *api) initRoutes() {
	r := mux.NewRouter()
	// save router to server
	a.srv.Handler = r
	
	// declare routes
	r.HandleFunc("/health", a.health).Methods("GET")
	r.HandleFunc("/api/v1/thing", a.getAllThings).Methods("GET")
	r.HandleFunc("/api/v1/thing", a.createThing).Methods("POST")
	r.HandleFunc("/api/v1/thing/{id}", a.getThing).Methods("GET")
	r.HandleFunc("/api/v1/thing/{id}", a.updateThing).Methods("PUT")
	r.HandleFunc("/api/v1/thing/{id}", a.deleteThing).Methods("DELETE")
}

