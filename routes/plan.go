package routes

import (
	"assignment8/controllers"
	"net/http"
)

func Plan(w http.ResponseWriter, r *http.Request) {
	controllers.Plan(w, r)
}
