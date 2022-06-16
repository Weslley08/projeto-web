CREATE DATABASE n2-web;

CREATE TABLE funcionarios (
	id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
	nome varchar(30) NOT NULL,
	sobrenome varchar(50) NOT NULL,
	cpf varchar(15) NOT NULL,
	datanascimento varchar NOT NULL,
	salario varchar(100) NOT NULL,
	cargo varchar NOT NULL,
	CONSTRAINT funcionarios_pkey PRIMARY KEY (id)
);