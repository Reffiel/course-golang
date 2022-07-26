package main

import (
	"database/sql"
	"fmt"
	"log"

	/*import implicito por causa do driver mysql, este vai utilizar este import para localizar no database/sql*/
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*Padrao da sctring para conexão "usuario:senha/banco", mais configurações adicionais como ?charset utf8 devido reconhecimento de caracteres e parsetime para reconhecer datas e local para reconhcer local */
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTimeTrue&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		fmt.Println("Dentro do SQL.Open")
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		fmt.Println("Dentro do ping")
		log.Fatal(erro)
	}
	fmt.Println("Conexão está aberta")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
