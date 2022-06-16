package Rotas

import (
	"projeto-web/controllers"
	"net/http"
)

func CarregarRotas() {
	http.HandleFunc("/", controllers.FindAll)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
}