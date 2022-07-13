package main

import "fmt"

/*Função variática base*/
func soma(numeros ...int) {
	fmt.Println(numeros)
}

/*Função variática soma*/
func soma2(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

/*Função concat exemplo*/
func escrever(txt string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(txt, numero)
	}
}

func main() {
	soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	totalSoma := soma2(10, 15, 5)
	fmt.Println(totalSoma)
	escrever("Hello!", 10, 2, 3, 4, 5, 10, 10000)
}

/*Variáticos somente é possível 1 por função e tem de ser obrigatoriamente o último*/
