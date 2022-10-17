package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	/*CRUD
	Create > POST
	Read > GET
	Update > PUT
	Delete > DELETE
	*/

	/*Criar um router onde vai conter todas as rotas para utilização
	Boa prática utilizar mais forma verbal e/ou singular
	*/
	router := mux.NewRouter()
	router.HandleFunc("/usuario", servidor.CriarUsuario).Methods(http.MethodPost)          //Constante contendo a string POST
	router.HandleFunc("/usuario", servidor.BuscarUsuarios).Methods(http.MethodGet)         //Constante contendo a string GET
	router.HandleFunc("/usuario/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)     //Constante contendo a string GET
	router.HandleFunc("/usuario/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)  //Constante contendo a string PUT
	router.HandleFunc("/usuario/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete) //Constante contendo a string DELETE

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}
