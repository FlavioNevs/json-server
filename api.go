package main

import (
	"github.com/gorilla/mux"
)

type Api struct {
	EndpointList []Endpoint `json:"endpoints"`
}

type Base struct {
	Api  Api                                 `json:"api"`
	Data map[string][]map[string]interface{} `json:"data"`
}

func (api *Api) CreateRouter() *mux.Router {
	router := mux.NewRouter()
	for _, endpoint := range api.EndpointList {
		endpoint.CreateEndpoint(router)
	}
	return router
}
