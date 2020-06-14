package main

import (
	"assignment8/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// os.Setenv("SPOON_API_KEY", SPOON_API_KEY)
	router := NewRouter()
	apiRoute := router.PathPrefix("/assignment").Subrouter()

	apiRoute.HandleFunc("/planner", routes.Plan).Methods("GET")
	fmt.Println("Running server on 8200")
	http.ListenAndServe(":8200", router)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}
