package models

import (
	"Go_web_app__/db"
)

// model para se relacionar com o db
type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// funçao para fazer o query no db
func GetProdutos() []Produto {
	db := db.Con()
	selectDeTodosProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	// loop para pegar todos os produtos
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
	// fechando a conexao do db
	defer db.Close()
	return produtos
}

// funçao para criar novo produto no banco
func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.Con()

	insereDadosDb, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insereDadosDb.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}
