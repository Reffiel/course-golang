package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}
func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários."))
}

func main() {
	/*HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB*/

	/*CLIENTE (FAZ A REQUISIÇÃO) - SERVIDOR (PROCESSA A REQUISIÇÃO E ENVIA RESPOSTA)*/

	// REQUEST - RESPONSIVE

	//ROTAS
	//URI - IDENTIFICADOR DO RECURSO (Identificar o caminho que deseja fazer)
	//MÉTODO - O QUE FAZER COM O RECURSO > GET, POST, PUT, DELETE

	/*Essas rotas somente escreve no ecrã*/
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)
	/*Aqui temos um servidor rodando nesta porta 5000, e o segundo parametro será visto nas proximas aulas*/
	log.Fatal(http.ListenAndServe(":5000", nil))
}
