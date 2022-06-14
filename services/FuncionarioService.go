package FuncionarioService

import (
	"projeto-web/models/entity"
	"projeto-web/db"
)

func FindAll() []FuncionarioStruct.Funcionario {
	db := database.ConectaBancoDados()
	buscaFuncionario, err := db.Query("SELECT * FROM funcionarios")
	if err != nil {
		panic(err.Error())
	}
	funcionario := FuncionarioStruct.Funcionario{}
	funcionarios := []FuncionarioStruct.Funcionario{}
	for buscaFuncionario.Next() {
		var Id int
		var Nome string
		var Sobrenome string
		var Cpf string
		var DataNascimento string
		var Salario string
		var Cargo string
		err = buscaFuncionario.Scan(&Id, &Nome, &Sobrenome, &Cpf, &DataNascimento, &Salario, &Cargo)
		if err != nil {
			panic(err.Error())
		}
		funcionario.Id = Id
		funcionario.Nome = Nome
		funcionario.Sobrenome = Sobrenome
		funcionario.Cpf = Cpf
		funcionario.DataNascimento = DataNascimento
		funcionario.Salario = Salario
		funcionario.Cargo = Cargo
		funcionarios = append(funcionarios, funcionario)
	}
	defer db.Close()
	return funcionarios
}