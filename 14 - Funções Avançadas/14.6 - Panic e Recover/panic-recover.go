package main

import "fmt"

/*recover utilizado com if para verificar se ha algo a recuperar assim é salvo a execução do programa*/
func recuperarExecucao() {
	fmt.Println("Entrou na recuperação da execução!")
	if r := recover(); r != nil {
		fmt.Println("Execução Recuperada com sucesso!")
		fmt.Println(r)
	}
}
func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	/*GO entra em modo panico e caso não ser utilizado o recover
	primeiro é chamado todos DEFER e após é finalizado o programa
	*/
	panic("A MÉDIA É IGUAL A 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução")
}
