package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// variavel que carrega todos  os template
var temp = template.Must(template.ParseGlob("templates/*.html"))

// funcao para receber a request da rota  e criar  o server
func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9000", nil)
}

// função para renderizar o template apontando para o http.Request
func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Verde,cor de lodo", Preco: 45, Quantidade: 20},
		{"Bermuda", "Cinza, claro", 35, 10},
		{"Sapato", "Preto,couro de Jacaré", 105, 8},
		{"Calça", "Azul,jeans escuro", 90, 5},
	}

	temp.ExecuteTemplate(w, "index", produtos)
}

// aula 3 terminada
