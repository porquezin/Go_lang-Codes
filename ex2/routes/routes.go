package routes

import (
	"ex2/controllers"
	"net/http"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
