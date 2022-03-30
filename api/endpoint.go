package api

import (
	"encoding/json"
	"fmt"
	"json-server/database"
	"log"
	"net/http"
)

type Endpoint struct {
	Name   string
	Fields []string
	Db     *database.Database
}

func EndpointFactory(name string, db *database.Database) Endpoint {
	var fields []string
	return Endpoint{name, fields, db}
}

func (e Endpoint) GetMethods() map[string]func(w http.ResponseWriter, r *http.Request) {
	return map[string]func(w http.ResponseWriter, r *http.Request){
		"Get":  e.Get,
		"Post": e.Post,
	}
}

func logRequests(method, uri, status_code string, ContentLength int64) {
	log.Printf("\"%s %s HTTP/1.1\" %s %d \n", method, uri, status_code, ContentLength)
}

func (e Endpoint) Get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(e.Db.Data[e.Name])

	logRequests(r.Method, r.RequestURI, "200", r.ContentLength)
}

func (e Endpoint) Post(w http.ResponseWriter, r *http.Request) {
	var d map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Println(err)
	}

	e.Db.Data[e.Name] = append(e.Db.Data[e.Name], d)
	json.NewEncoder(w).Encode(d)

	logRequests(r.Method, r.RequestURI, "200", r.ContentLength)
}
