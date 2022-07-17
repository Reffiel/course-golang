package main

import "fmt"

/*Métodos tem uma estrutura para utilizar*/
type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)

}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

/*Comum utilizar ponteiros para alterar em banco de dados*/
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()
	maiorIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorIdade)

	usuario2 := usuario{"Davi", 5}
	fmt.Println(usuario2)
	usuario2.salvar()
	maiorIdade2 := usuario2.maiorDeIdade()
	fmt.Println(maiorIdade2)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
