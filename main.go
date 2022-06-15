package main

import (
	"projeto-web/routes"
	"net/http"
	"fmt"
)

func main() {
	Rotas.CarregarRotas()
	http.ListenAndServe(":8000", nil)
	fmt.Println("Servidor iniciado na porta 8000")
}