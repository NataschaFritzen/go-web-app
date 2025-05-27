package db

import (
	"database/sql"

	_ "github.com/lib/pq" // Motivo do "_", para utilizar durante o tempo de execução da aplicação
)

// Insira suas credenciais de acesso ao banco de dados
func ConectaComBancoDeDados() *sql.DB {
	conexao := "user= dbname= password= host= sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
