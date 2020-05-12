package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"net/http"
)

// common response object
type Resp struct {
	Body string `json:"body"`
}

// request object for POST
type CreateReq struct {
	Name     string   `json:"name"`
	Features []string `json:"features"`
}

const thingTableName = "things"

// handles creating thing upon POST request with json data
func (a *api) createThing(w http.ResponseWriter, r *http.Request) {
	var in CreateReq
	
	// decode request body json to our object
	_ = json.NewDecoder(r.Body).Decode(&in)
	
	// build sql query
	featuresAggregate := "'{"
	for _, v := range in.Features {
		featuresAggregate += v + ","
	}
	
	featuresAggregate = featuresAggregate[0 : len(featuresAggregate)-1]
	featuresAggregate += "}'"
	
	// prepare query
	stmt := fmt.Sprintf("INSERT INTO %s (name, features) VALUES ('%s', %s)", thingTableName, in.Name, featuresAggregate)
	
	// execute query
	_, err := a.db.Exec(stmt)
	if err != nil {
		out := Resp{
			Body: err.Error(),
		}
		
		jsonResponse(w, out, 500)
		return
	}
	
	out := Resp{
		Body: "created",
	}
	
	// return json response
	jsonResponse(w, out, 200)
}

// get handler for getting a thing by id upon GET request
func (a *api) getThing(w http.ResponseWriter, r *http.Request) {
	// path variables {id}
	vars := mux.Vars(r)
	
	// prepare query
	stmt := fmt.Sprintf("SELECT * FROM %s WHERE id=%s", thingTableName, vars["id"])
	
	// execute query
	rows, err := a.db.Query(stmt)
	if err != nil {
		out := Resp{
			Body:   err.Error(),
		}
		
		jsonResponse(w, out, 500)
		return
	}
	
	// build response from sql result rows
	var out Thing
	for rows.Next() {
		err := rows.Scan(&out.ID, &out.Name, (*pq.StringArray)(&out.Features))
		if err != nil {
			out := Resp{
				Body:   err.Error(),
			}
			
			jsonResponse(w, out, 500)
			return
		}
	}
	
	// check for null data
	if out.ID == 0 {
		out := Resp{
			Body: "no data found",
		}
		
		jsonResponse(w, out, 404)
		return
	}
	
	jsonResponse(w, out, 200)
}

// get all handler for getting all things upon GET request
func (a *api) getAllThings(w http.ResponseWriter, r *http.Request) {
	// prepare query
	stmt := fmt.Sprintf("SELECT * FROM %s", thingTableName)
	
	// execute query
	rows, err := a.db.Query(stmt)
	if err != nil {
		out := Resp{
			Body:   err.Error(),
		}
		
		jsonResponse(w, out, 500)
		return
	}
	
	// build response from sql result rows
	var out []Thing
	for rows.Next() {
		var t Thing
		err := rows.Scan(&t.ID, &t.Name, (*pq.StringArray)(&t.Features))
		if err != nil {
			out := Resp{
				Body:   err.Error(),
			}
			
			jsonResponse(w, out, 500)
			return
		}
		
		out = append(out, t)
	}
	
	// check for null data
	if len(out) == 0 {
		out := Resp{
			Body: "no data found",
		}
		
		jsonResponse(w, out, 404)
		return
	}
	
	jsonResponse(w, out, 200)
}

// update handler for updating a thing by id upon PUT request
func (a *api) updateThing(w http.ResponseWriter, r *http.Request) {
	var in CreateReq
	// decode request body json to our object
	_ = json.NewDecoder(r.Body).Decode(&in)
	// path variables {id}
	vars := mux.Vars(r)
	// prepare query
	stmt := fmt.Sprintf("UPDATE %s SET name='%s' WHERE id=%s", thingTableName, in.Name, vars["id"])
	// execute query
	_, err := a.db.Exec(stmt)
	if err != nil {
		out := Resp{
			Body:   err.Error(),
		}
		
		jsonResponse(w, out, 500)
		return
	}
	
	// return json response
	out := Resp{
		Body: "updated",
	}
	
	jsonResponse(w, out, 200)
}

// delete handler for deleting a thing by id upon DELETE request
func (a *api) deleteThing(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=%s", thingTableName, vars["id"])
	
	_, err := a.db.Exec(stmt)
	if err != nil {
		out := Resp{
			Body:   err.Error(),
		}
		
		jsonResponse(w, out, 500)
		return
	}
	
	out := Resp{
		Body: "deleted",
	}
	
	jsonResponse(w, out, 200)
}

// health endpoint
func (a *api) health(w http.ResponseWriter, _ *http.Request) {
	out := Resp{
		Body: "I'm Ok :)",
	}
	
	jsonResponse(w, out, 200)
}

// helper function for returning json response
func jsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
