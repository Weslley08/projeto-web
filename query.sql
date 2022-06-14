PASSO 1: create database n2-web;

PASSO 2: 
create table funcionarios (
id int primary key not null generated always as identity,
nome varchar(30) not null,
sobrenome varchar(50) not null,
cpf varchar(15) not null,
dataNascimento varchar not null,
salario varchar(100) not null,
cargo varchar not null
)

PASSO 3: 
INSERT INTO funcionarios  (nome, sobrenome, cpf, datanascimento, salario, cargo) 
VALUES ('Weslley Jonathan', 'Cezario Santana da Silva', '12345678900', '08/12/2000', '1.000', 'Engenheiro de software pleno');
