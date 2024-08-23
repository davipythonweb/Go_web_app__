# Go_web_app__

### Criando uma Aplicação Web- MVC com Go

- banco de dados
`postgres`

* fazendo o build do projeto
 `go build .`

* buildando e rodando todos os arquivos
`go run *.go`

- instalar pacote Go, do git:
`go install github.com/lib/pq@v1.10.9`

- codigo sql para criar a tabela no banco:

"""
create table produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
)
"""

- codigo sql para inserir um produto no banco:

"""
insert into produtos (nome, descricao, preco, quantidade) values
('camisa', 'vinho', 23.50, 10),
('Bone', 'preto', 65.75, 3);
"""

- fazendo um filtro sql para mostrar as 100 primeiras linhas da tabela:

"""
SELECT * FROM public.produtos
ORDER BY id ASC LIMIT 100

"""