package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq" // Motivo do "_", para utilizar durante o tempo de execução da aplicação
)

// Insira suas credenciais de acesso ao banco de dados
func conectaComBancoDeDados() *sql.DB {
	conexao := "user dbname password host sslmode"

}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Mouse", "Rápido", 15, 3},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
