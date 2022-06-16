package FuncionarioService

import (
	database "projeto-web/db"
	FuncionarioStruct "projeto-web/models/entity"
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

func CriaNovoFuncionario(nome string, sobrenome string, cpf, datanascimento string, salario string, cargo string) {
	db := database.ConectaBancoDados()

	insereDadosNoBanco, err := db.Prepare("insert into funcionarios(nome, sobrenome, cpf, datanascimento, salario, cargo) values($1, $2, $3, $4, $5, $6)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, sobrenome, cpf, datanascimento, salario, cargo)
	defer db.Close()
}

func EditaFuncionario(id string) FuncionarioStruct.Funcionario {
	db := database.ConectaBancoDados()
	funcionarioDoBanco, err := db.Query("select * from funcionarios where id=$1", id)
	if err != nil {
		panic(err.Error())
	}
	funcionarioParaAtualizar := FuncionarioStruct.Funcionario{}

	for funcionarioDoBanco.Next() {
		var id int
		var nome, sobrenome string
		var cpf, DataNascimento string
		var salario, cargo string

		err = funcionarioDoBanco.Scan(&id, &nome, &sobrenome, &cpf, &DataNascimento, &salario, &cargo)
		if err != nil {
			panic(err.Error())
		}
		funcionarioParaAtualizar.Id = id
		funcionarioParaAtualizar.Nome = nome
		funcionarioParaAtualizar.Sobrenome = sobrenome
		funcionarioParaAtualizar.Cpf = cpf
		funcionarioParaAtualizar.DataNascimento = DataNascimento
		funcionarioParaAtualizar.Salario = salario
		funcionarioParaAtualizar.Cargo = cargo
	}
	defer db.Close()
	return funcionarioParaAtualizar
}

func AtualizaFuncionario(id int, nome string, sobrenome string, cpf string, datanascimento string, salario string, cargo string) {
	db := database.ConectaBancoDados()

	AtualizaFuncionario, err := db.Prepare("update funcionarios set nome=$1, sobrenome=$2, cpf=$3, datanascimento=$4, salario=$5, cargo=$6 where id=$7")
	if err != nil {
		panic(err.Error())
	}
	AtualizaFuncionario.Exec(nome, sobrenome, cpf, datanascimento, salario, cargo, id)
	defer db.Close()
}

func DeletaFuncionario(id string) {
	db := database.ConectaBancoDados()

	deletarFuncionario, err := db.Prepare("delete from funcionarios where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarFuncionario.Exec(id)
	defer db.Close()
}
