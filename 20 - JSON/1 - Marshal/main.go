package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json: "nome"`
	Raca  string `json: raca"`
	Idade uint   `json: idade"`
}

func main() {

	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)
	/*
		json.Marshal()
		Utilizado para converter um map para json ou estruct para json

		json.Unmarshal()
		Utilizado para converter um json para map ou json para estruct
	*/
	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	/*desta forma somente temos um array de bytes, slices do tipo uint8 > [123 34 78 111 109 101 34...*/
	fmt.Println(cachorroEmJson)

	/*Para visualizar precisa utilizar outro pacote do go bytes*/
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	/*Para inserir do tipo map*/
	c2 := map[string]string{
		"nome":  "Toby",
		"idade": "Poodle",
	}
	cachorroEmJson2, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)

	}
	fmt.Println(cachorroEmJson2)
	fmt.Println(bytes.NewBuffer(cachorroEmJson2))

}
