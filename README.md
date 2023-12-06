# <div align="center">Piravraria<div>

## Equipe

- [Guilherme Farias (2214635)](https://github.com/G-FARIAS-71)
- [Lucca Han (2214677)](https://github.com/Ludihan)
- [Matheus Najal (2219156)](https://github.com/matheusnajal)
- [Marcos Antônio (2214668)](https://github.com/jekker9)

## Visão geral do Repositório

Este é o repositório do trabalho final da disciplina Ambiente de dados da Unifor, a qual ensina a usar banco de dados.

O projeto foi feito usando o SGBD open-source postgress e a linguagem de programação Go e segue as orientações fornecidas [aqui](orientacoes.pdf).

O banco de dados foi feito para uma livraria fictícia chamada Piravraria. Mais informações sobre ela podem ser encontradas [aqui](piravraria.pdf).

## Orientações

### Passo 1

> Criar o Diagrama Entidade Relacionamento (DER) e o modelo lógico do Diagrama Entidade Relacionamento (MER)

[Clique aqui para ver os diagramas](https://miro.com/app/board/uXjVNMDUHlk=/)

### Passo 2

> Com base no modelo lógico criar o banco de dados no MySQL e fazer a entrega dos scripts de criação do banco de dados em um arquivo texto separado (0.5)

[Clique aqui para ver o script](piravraria.sql)

### Passo 3

> Utilizando linguágem de programa, faça um programa que realiza o CRUD* de pelo menos 1 entidade do seu Diagrama Entidade Relacionamento. Você deve também:

1. [Realizar uma consulta utilizando JOIN](src/cli/table.go)

2. [Utilizar transação nos INSERTS](README.md)

3. [Criar um método para cadastrar em massa dados provenientes de um arquivo CSV.](README.md) Obrigatoriamente os INSERTS realizados nesta atividades devem ser inseridos no banco através de uma única transação

4. [Realizar uma exportação para CSV de pelo menos uma consulta executada no banco.](README.md)

### Passo 4

[Criar índices nas tabelas mais utilizadas](README.md)

### Passo 5

[Entregar em um documento com consultas (em SQL e algebra relacional) definidas para a problemática de cada grupo](consultas.md)