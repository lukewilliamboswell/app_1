package api

import (
	"database/sql"
	"log"
	"net/http"
)

// api holds data for running api server
type api struct {
	db  *sql.DB
	srv *http.Server
}

// api constructor function
func NewApi(db *sql.DB) (*api, error) {
	newApi := api{srv: new(http.Server), db: db}
	// initialize routes
	newApi.initRoutes()

	// create new table upon start (will not create if one already exists)
	// _, err := newApi.db.Exec("CREATE TABLE IF NOT EXISTS things(id serial PRIMARY KEY, name text, features text[])")
	// if err != nil {
	// 	return nil, err
	// }

	return &newApi, nil
}

// function for running api
func (a *api) Run(addr string) error {
	// saving the address to server
	a.srv.Addr = addr

	log.Printf("running api server at http://0.0.0.0%s\n", addr)
	return a.srv.ListenAndServe()
}
