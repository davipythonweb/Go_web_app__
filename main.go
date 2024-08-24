package main

import (
	"html/template"
	"net/http"

	"Go_web_app__/models"
)

// variavel que carrega todos  os template
var temp = template.Must(template.ParseGlob("templates/*.html"))

// funcao para receber a request da rota  e criar  o server
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9100", nil)
}

// função para renderizar o template apontando para o http.Request
func index(w http.ResponseWriter, r *http.Request) {

	// pega todos os produtos do banco e armazena na variavel
	allProdutos := models.GetProdutos()
	// chamando o template e a variavel com todos os produtos do db
	temp.ExecuteTemplate(w, "index", allProdutos)

}

// aula 7 terminada
