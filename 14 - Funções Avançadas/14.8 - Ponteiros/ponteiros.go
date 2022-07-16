package main

import "fmt"

/*Desta forma estamos a utilizar um cópia do valor, não altera o valor*/
func inverterSinal(numero int) int {
	return numero * -1
}

/*E com o ponteiro estamos a utilizar a referência da variavel e assim alteramos o valor*/
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	numeroNeg := -10
	numeroInvertido = inverterSinal(numeroNeg)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println("Valor inicial da variavel-> ", novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println("Variável alterada -> ", novoNumero)
}
