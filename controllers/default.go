package controllers

import (
	"Go_web_app__/models"
	"net/http"
	"text/template"
)

// variavel que carrega todos  os template
var temp = template.Must(template.ParseGlob("templates/*.html"))

// função para renderizar o template apontando para o http.Request
func Index(w http.ResponseWriter, r *http.Request) {

	// pega todos os produtos do banco e armazena na variavel
	allProdutos := models.GetProdutos()
	// chamando o template e a variavel com todos os produtos do db
	temp.ExecuteTemplate(w, "Index", allProdutos)

}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
