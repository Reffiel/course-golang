package main

import "fmt"

/*Função closure sempre busca as variaveis da referencia inicial, sempre guarda a sua referencia inicial*/
func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da func main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()

}
