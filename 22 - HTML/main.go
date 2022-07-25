package main

import (
	"fmt"
	"html/template" //Criar templates baseados no arquivo html e deixar dinâmico
	"log"
	"net/http"
)

/*
Pra seguir a convenção deve ser fora da função main
Não é um Slice mas tem mais de uma template neste tipo devido a estrutura de arvore
*/
var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	/*Aqui esta a passar todos html na mesma pasta para templates*/
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{"João", "joao.pedro@gmail.com"}
		/*Para esta chamada são necessários 3 parametros	1-ResponseWriter 2-O nome do template 3-Se tivesse um dado para passar para o template(como nome de usuário, um dado, etc..) para tornar mais dinâmica*/
		templates.ExecuteTemplate(w, "home.html", u)
	})
	/* Abaixo exemplo de API simples retorna mensagem "Olá mundo!"
	http.HandleFunc("/home", func(w http.ResponseWriter, r *httr.Request) {
		w.Write([]byte("Olá mundo!"))
	})
	*/
	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
