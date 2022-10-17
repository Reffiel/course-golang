package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

/*CriarUsuario insere um usuário no banco de dados
Convenção utilizar ao inves de erro o "err"

Validar retorno negativo de todas verificações
Fazer função para response writer

Pacote de log para acompanhar erros de execução
*/
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(erro.Error()))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (? , ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao carregar dados do usuario"))
		return
	}
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id inserido"))
		return
	}
	/*STATUS CODES > Ex: 201, 404, etc no go pode ser utilizado a constante*/
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso no id: %d", idInserido)))
}

//BuscarUsuarios busca todos usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o Banco de dados!"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		/*importante a ordem das colunas conforme o BD*/
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuario!"))
			return

		}
		usuarios = append(usuarios, usuario)

	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter usuario para o json!"))
		return

	}

}

//BuscarUsuario efetua busca de um usuário específico no bando de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuário"))
		return
	}
	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuario"))
			return
		}
	}

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuário não encontrado"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario para JSON"))
		return

	}

}

//AtualizarUsuario altera dados de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler corpo da requisição"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para sctruct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário!"))
		return
	}
	/*Para atualização de um específico mais vale 200*/
	w.WriteHeader(http.StatusNoContent)

}

//Deletar usuário remove um usuário na base de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	result := teste("teste")
	fmt.Println(result)

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar usuario"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func teste(parm string) string {
	if parm == "rodrigo" {
		return "sim"
	}
	return ""
}
