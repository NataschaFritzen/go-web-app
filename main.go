package main

import (
	"go-web-app/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":3000", nil)
}
