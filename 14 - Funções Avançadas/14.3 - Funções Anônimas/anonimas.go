package main

import "fmt"

func main() {
	/*Função anônimas*/
	func() {
		fmt.Println("Olá mundo!")
	}()

	/*Pode ter parâmetro*/
	func(txt string) {
		fmt.Println(txt)
	}("Parâmetro")

	/*Pode ter retorno*/
	retorno := func(texto string) string {
		/*Função Sprintf concatena String com outros tipos de dados
		 	%s é para tipo string
			%d é para tipo número
		*/
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Retorno")
	fmt.Println(retorno)
}
