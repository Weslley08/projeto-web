package Rotas

import (
	"projeto-web/controllers"
	"net/http"
)

func CarregarRotas() {
	http.HandleFunc("/", controllers.FindAll)
//	http.HandleFunc("/id", controllers.FindByCpf)
//	http.HandleFunc("/create", controllers.Create)
//	http.HandleFunc("/update", controllers.Update)
//	http.HandleFunc("/delete", controllers.Delete)
}