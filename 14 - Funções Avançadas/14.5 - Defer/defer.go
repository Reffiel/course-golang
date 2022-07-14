package main

import "fmt"

func funcao1() {
	fmt.Println("Exec func 1")
}
func funcao2() {
	fmt.Println("Exec func 2")
}

/*em uma função é executada a linha do defer imediatamente antes do return
normalmente utilizada para fechar conexão quando a função faz conexão com o banco de dados*/
func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Inicio da função para verificar se aluno esta aprovado")
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	}
	return false
}

func main() {
	/*DEFER = Adiar*/
	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
