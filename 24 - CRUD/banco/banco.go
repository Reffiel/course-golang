package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de conexão com MySQL
)

/*Como não estamos dentro do arquivo main não se pode fechar com fatal temos de retornar algo, neste caso retornaremos ou o ponteiro de conexão ou o erro, isso porque são mutuamente exclusivos*/
func Conectar() (*sql.DB, error) {
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
