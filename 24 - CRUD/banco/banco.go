package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão com MySQL
)

/*Como não estamos dentro do arquivo main não se pode fechar com fatal temos de retornar algo, neste caso retornaremos ou o ponteiro de conexão ou o erro, isso porque são mutuamente exclusivos
Exemplo de interface para conectar uma vez, na inicialização faz-se o contrato para conexão. Com a instância já se consegue utilizar todos handle
para fazer downliad di driver my sql comando > go get github.com/go-sql-driver/mysql
*/
func Conectar() (*sql.DB, error) {
	/*String para conexão para mysql > usuario:senha@/nomeDoBanco?charset=utf8&parseTime=True&loc=Local
	| postgres é diferente*/
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTimeTrue&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}
	/*Defer não é feito aqui neste caso, e sim nas funções*/
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
