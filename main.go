package main

import (
	"html/template"
	"net/http"
)

// variavel que carrega o template
var temp = template.Must(template.ParseGlob("templates/*.html"))

// funcao para receber a request da rota  e criar  o server
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9000", nil)
}

// função para renderizar o template apontando para o http.Request
func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index", nil)
}

// aula 3 terminada
