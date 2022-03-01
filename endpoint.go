package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Endpoint struct {
	Name  string `json:"name"`
	Route string `json:"route"`

	AllowedMethods []string `json:"allowed_methods"`
}

func (e *Endpoint) CreateEndpoint(route *mux.Router) {
	for _, method := range e.AllowedMethods {
		route.HandleFunc(
			e.GetRoute(),
			e.GetMethod(method),
		).Methods(method)
	}

}
func (e Endpoint) GetRoute() string {
	if e.Route != "" {
		return e.Route
	}
	return fmt.Sprintf("/%s", strings.ToLower(e.Name))
}

func (e Endpoint) GetMethod(method string) func(http.ResponseWriter, *http.Request) {
	switch method {
	case "GET":
		return e.Get
	case "POST":
		return e.Post
	default:
		return func(http.ResponseWriter, *http.Request) {}
	}
}

func (e Endpoint) Get(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(database[e.Name])
}

func (e Endpoint) Post(w http.ResponseWriter, r *http.Request) {
	var d map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Println(err)
	}

	database[e.Name] = append(database[e.Name], d)
	json.NewEncoder(w).Encode(d)
}
