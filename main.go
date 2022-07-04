package main

import (
	"fmt"
	"modulo/auxiliar"
)

/*
go guild > biuld precisa ser atualizado sempre que for necessário
go install faz o mesmo que go build, porém adiciona na pasta raiz do projeto para execução
*/

func main() {
	fmt.Println("Writing main archive")
	auxiliar.Escrever()
}
