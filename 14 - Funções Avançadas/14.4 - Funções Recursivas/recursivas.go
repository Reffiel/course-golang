package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	//Fibonacci 1 1 2 3 5 8 13 21

	posicao := uint(1)
	fmt.Println(fibonacci(posicao))

	posicao2 := uint(12)
	for i := uint(1); i < posicao2; i++ {
		//fmt.Println(fibonacci(i))
		fmt.Println(fmt.Sprintf("Posição %d fibonacci é igual a %d", i, fibonacci(i)))
	}

}
