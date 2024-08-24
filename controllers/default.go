package controllers

import (
	"Go_web_app__/models"
	"log"
	"net/http"
	"strconv"
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

// funçao para inserir os dados
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do Preço:", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}

	http.Redirect(w, r, "/", 301)
}

// função para deletar pelo id
func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

// função para editar pelo id
func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditarProduto(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}
