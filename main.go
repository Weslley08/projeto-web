package main

import (
	"projeto-web/routes"
	"net/http"
	"fmt"
)

func main() {
	Rotas.CarregarRotas()
	fmt.Println("Servidor iniciado na porta 8000")
	http.ListenAndServe(":8000", nil)
}