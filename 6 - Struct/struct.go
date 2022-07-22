package main

import "fmt"

type usuario struct {
	Nome     string
	Idade    uint8
	Endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")
	/*Quando não é populado é atribuído sempre o valor 0 de cada tipo que contem no struct*/
	var u usuario
	u.Nome = "Davi"
	u.Idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Via de Matti", 0}

	usuario2 := usuario{"Davi", 21, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{Nome: "Davi"}
	fmt.Println(usuario3)
}
