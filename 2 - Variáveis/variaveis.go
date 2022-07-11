package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1 Texto"
	/*Declaração implicita de string, inferência de tipo, inferindo tipo baseado no valor dele*/
	variavel2 := "Variável 2 Texto"
	fmt.Println(variavel1 + "**\"*" + variavel2)

	var (
		variavel3 string = "Tipo de instanciação 3"
		variavel4 string = "Tipo de instanciação 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	/*Troca de valores entre variáveis*/
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
