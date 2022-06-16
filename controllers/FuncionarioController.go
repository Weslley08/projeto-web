package controllers

import (
	"html/template"
	"log"
	"net/http"
	"projeto-web/services"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func FindAll(w http.ResponseWriter, r *http.Request) {
	findAll := FuncionarioService.FindAll()
	templates.ExecuteTemplate(w, "Index", findAll)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		sobrenome := r.FormValue("sobrenome")
		cpf := r.FormValue("cpf")
		DataNascimento := r.FormValue("DataNascimento")
		salario := r.FormValue("salario")
		cargo := r.FormValue("cargo")

		FuncionarioService.CriaNovoFuncionario(nome, sobrenome, cpf, DataNascimento, salario, cargo)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idFuncionario := r.URL.Query().Get("id")
	funcionario := FuncionarioService.EditaFuncionario(idFuncionario)
	templates.ExecuteTemplate(w, "Edit", funcionario)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		sobrenome := r.FormValue("sobrenome")
		cpf := r.FormValue("cpf")
		DataNascimento := r.FormValue("DataNascimento")
		salario := r.FormValue("salario")
		cargo := r.FormValue("cargo")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		FuncionarioService.AtualizaFuncionario(idConvertidaParaInt, nome, sobrenome, cpf, DataNascimento, salario, cargo)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}


func Delete(w http.ResponseWriter, r *http.Request) {
	idFuncionario := r.URL.Query().Get("id")
	FuncionarioService.DeletaFuncionario(idFuncionario)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
