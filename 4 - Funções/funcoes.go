package main

import "fmt"

/*Tipo de retorno após as variáveis*/
func somar(n1 int8, n2 int8) int8 {
	return (n1 + n2)
}

/*Retorno de dois valores*/
func somar2(n3 int8, n4 int8) (int8, int8) {
	return (n3 + n4), (n3 - n4)
}

func main() {
	soma := somar(3, 12)
	fmt.Println(soma)

	soma2, subtracao := somar2(15, 2)
	fmt.Println(soma2, subtracao)

	/*Para retornar somente um valor em funções que retornam mais de um utilizar "_"*/
	_, subtracao2 := somar2(15, 2)
	fmt.Println(subtracao2)
	soma3, _ := somar2(15, 2)
	fmt.Println(soma3)

	/*Variável do tipo função*/
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Função f")
	fmt.Println(resultado)

	/*Função é um tipo, atribuir a variavel, pode passar por parametro e pode retornar uma função também*/

}
