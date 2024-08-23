package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func con() *sql.DB {

	conexao := "user=postgres dbname=golang_storage password=Db5$Ades10 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db

}

type Produto struct {
	Id         int
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
	http.ListenAndServe(":9100", nil)
}

// função para renderizar o template apontando para o http.Request
func index(w http.ResponseWriter, r *http.Request) {

	db := con()
	selectDeTodosProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "index", produtos)

	defer db.Close()

}

// aula 7 terminada
