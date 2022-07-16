package main

import "fmt"

func main() {
	fmt.Println("Função main sendo executada")
}

/*
Init é uma função que é executada antes da main, indiferente do ponto onde esteja no ficheiro,
Diferemte da função main que é possivel somente uma, init pode ter uma por ficheiro.
Pode ser utilizada para função de um setup, inicialazação de varíaveis
*/
func init() {
	fmt.Println("Executando a função init")
}
