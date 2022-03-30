package api

import "github.com/gorilla/mux"

func (api *Api) RouterFactory() *mux.Router {
	router := mux.NewRouter()
	for _, end := range api.endpoints {
		for method, fc := range end.GetMethods() {
			router.HandleFunc("/"+end.Name, fc).Methods(method)
		}
	}

	return router
}
