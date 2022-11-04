package controllers

import (
	"ex2/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProd := models.BuscaAll()
	temp.ExecuteTemplate(w, "Index", allProd)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		desc := r.FormValue("descricao")

		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
		if err != nil {
			log.Panicln(err)
		}

		quant, err := strconv.Atoi(r.FormValue("quantidade"))
		if err != nil {
			log.Panicln(err)
		}

		models.AddProd(nome, desc, preco, quant)

		http.Redirect(w, r, "/", 301)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	models.Delete(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	prod := models.Edit(id)
	temp.ExecuteTemplate(w, "Edit", prod)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		i := r.FormValue("id")
		nome := r.FormValue("nome")
		desc := r.FormValue("descricao")
		p := r.FormValue("preco")
		q := r.FormValue("quantidade")

		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err.Error())
		}
		preco, err := strconv.ParseFloat(p, 64)
		if err != nil {
			panic(err.Error())
		}
		quant, err := strconv.Atoi(q)
		if err != nil {
			panic(err.Error())
		}

		models.Update(id, nome, desc, preco, quant)
	}
	http.Redirect(w, r, "/", 301)
}
