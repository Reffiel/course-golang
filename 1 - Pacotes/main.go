package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

/*
go guild > biuld precisa ser atualizado sempre que for necessário
go install faz o mesmo que go build, porém adiciona na pasta raiz do projeto para execução
após build pode rodar somente com o nome do módulo (tem de dizer onde esta este ficheiro "./modulo")
go get <url> para importar pacote externo, executar no nivel de pasta dos modulos
go mod tidy > remove dependências não utilizadas
*/

func main() {
	fmt.Println("Writing main archive")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("testegmail.com")
	fmt.Println(erro)
}
