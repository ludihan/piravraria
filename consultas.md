# <div align="center">Consultas</div>

## 1. Selecione todos os livros (SQL e Álgebra Relacional)

### SQL

```sql
SELECT *
FROM livro;
```

### Álgebra relacional

$livro$

## 2. Selecione a quantidade de registros na tabela livros (SQL e Álgebra Relacional) 

### SQL

```sql
SELECT count(*)
FROM livro;
```

### Álgebra relacional

$ \gamma_{\text{count}()}(\text{livro}) $

## 3. Selecione o nome e e-mail do autor (SQL e Álgebra Relacional) 

### SQL

```sql
SELECT nome,
       email
FROM autor;
```

### Álgebra relacional

$ \pi_{\text{nome}, \text{email}}(\text{autor}) $

## 4. Selecione os autores que nasceram no mês de dezembro (SQL)

```sql
SELECT *
FROM autor
WHERE extract(MONTH
              FROM data_nascimento) = 12;
```

## 5. Selecione os livros que tem idioma pt e en (SQL)

```sql
SELECT *
FROM livro
WHERE idioma = 'pt'
  OR idioma = 'en';
```

## 6. Selecione o título dos livros e o nome da editora (SQL) 

```sql
SELECT livro.titulo,
       editora.nome
FROM livro
INNER JOIN editora ON livro.editora_id = editora.id;
```

## 7. Selecione o nome do autor e a quantidade de livros que ele escreveu (SQL) 

```sql
SELECT autor.nome,
       count(livro_autor.autor_id)
FROM autor
INNER JOIN livro_autor ON livro_autor.autor_id = autor.id
GROUP BY autor.nome,
         livro_autor.autor_id;
```

## 8. Delete todos os livros do autor de código 1 (SQL) 

```sql
BEGIN;

DELETE
FROM desconto USING livro_autor
WHERE desconto.isbn = livro_autor.livro_isbn
  AND livro_autor.autor_id = 1;

DELETE
FROM livro_autor
WHERE livro_autor.autor_id = 1;

DELETE
FROM livro USING livro_autor
WHERE livro.isbn = livro_autor.livro_isbn
  AND livro_autor.autor_id = 1;

COMMIT;
```

## 9. Incremente em 20% o preço de todos os livros (SQL) 10- Insira 3 autores novos em um insert só (SQL)

```sql
UPDATE livro
SET preco_base = preco_base * 1.20;
```

```sql
SELECT *
FROM livro;
```

## 10. Insira 3 autores novos em um insert só (SQL)

```sql
INSERT INTO autor (nome, pseudonimo, email, genero, data_nascimento, data_falecimento, nacionalidade)
VALUES ('Jorge','gigante','jorge@gmail.com','m','1970-12-12','2020-05-30','Mexicano'),
       ('Marcos','jekker','marcos@gmail.com','m','1990-12-12','2019-05-30','Brasileiro'),
       ('Lucca','careca','lucca@gmail.com','m','1950-12-12','2021-05-30','Brasileiro');
```