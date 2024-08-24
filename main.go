package main

import (
	"Go_web_app__/routes"
	"net/http"
)

// funcao para receber a request da rota  e criar  o server
func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":9100", nil)
}

// aula 7 terminada
