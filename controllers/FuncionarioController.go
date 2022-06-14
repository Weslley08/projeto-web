package controllers

import (
	"projeto-web/services"
	"net/http"
	"html/template"
) 

var templates = template.Must(template.ParseGlob("templates/*.html"))

func FindAll(w http.ResponseWriter, r *http.Request) {
	findAll := FuncionarioService.FindAll()
	templates.ExecuteTemplate(w, "index", findAll)
}