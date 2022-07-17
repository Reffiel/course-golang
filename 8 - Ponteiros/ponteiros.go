package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variable11 int = 10
	var variable12 int = variable11

	fmt.Println(variable11, variable12)

	variable11++
	fmt.Println(variable11, variable12)

	/*Ponteiro é uma referência de memoria*/
	var variable13 int
	/*Para indicar um ponteiro precisa adicionar o asterisco*/
	var ponteiro *int

	variable13 = 100
	/*Para atribuir o endereço de memória para a variável adicionar o & */
	ponteiro = &variable13

	fmt.Println(variable13, ponteiro)
	variable13 = 150
	/*Desreferenciação*/
	fmt.Println(variable13, *ponteiro)

}
