package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*Caso não se queira que apareça a tag por exemplo do nome devemos adicionar um - ex: "Nome  string `json: "-"`"*/
type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: raca"`
	Idade uint   `json: idade"`
}

func main() {
	cachorroEmJson := `{"nome":"Rex","raca":"Dálmata","idade":3}`
	/*Declaração do structy em branco melhor com var para ficar mais claro mas é possível por atribuição do tipo ex: "c := cachorro"*/
	var c cachorro
	fmt.Println(c)
	/*Para passar os dados para dentro do struct precisamos utilizar o Unmarshal,
	1- Para isso precisa de dois parametros: os dados para passar e o endereço onde vai ficar os dados(Tem de ser com & para ser o ponteiro)
	2- Unmarshal precisa receber um slice de bytes por isso devemos adicionar json.Unmarsha...[]byte(cachorroEmJson)...(desta forma estamos a converter uma string em slice de bytes)
	3- Unmarshal retorna só um valor que é um Erro, e precisamos tratar
	*/

	if erro := json.Unmarshal([]byte(cachorroEmJson), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Primeiro", c)

	cachorro2EmJson := `{"nome":"Toby","raca":"Poodle"}`
	/*É preciso ter cuidado ao utilizar o map, com referência ao tipo a conversão tem de ser entre tipos que são permitidos*/
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJson), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)

}
