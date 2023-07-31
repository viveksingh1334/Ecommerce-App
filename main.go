package main

import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, nil)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/products.html"))
	tmpl.Execute(w, products)
}

func cartHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/cart.html"))
	tmpl.Execute(w, nil)
}
