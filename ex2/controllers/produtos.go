package controllers

import (
	"ex2/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProd := models.BuscaAll()
	temp.ExecuteTemplate(w, "Index", allProd)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
