# Go_web_app__

### Criando uma Aplicação Web- MVC com Go

- CriarProduto, EditarProduto, DeletearProduto

- Bootstrap-v5 via CDN

- banco de dados
`postgres`

* fazendo o build do projeto
 `go build .`

* buildando e rodando todos os arquivos
`go run *.go`

- instalar pacote Go, do git:
`go install github.com/lib/pq@v1.10.9`
`go get github.com/lib/pq`
`go get github.com/joho/godotenv`

- iniciar pacotes go
`go mod init`



- codigo para carregar .env:

"""
import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
    fmt.Println("TEST:", os.Getenv("TEST"))
}
"""