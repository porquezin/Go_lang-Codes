package main

import (
	"ex2/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe("0.0.0.0:8000", nil)
}
