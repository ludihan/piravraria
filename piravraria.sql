CREATE database piravraria;

-- rode isso se estiver no psql
-- \c piravraria

CREATE TABLE editora (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    data_fundacao DATE NULL,
    royalties float CHECK (royalties BETWEEN 0 and 100) NOT NULL
);

CREATE TYPE classificacao_etaria AS ENUM ('l', '10', '12', '14', '16', '18');

CREATE TABLE livro (
    isbn VARCHAR(13) PRIMARY KEY,
    editora_id INT NULL REFERENCES editora(id),
    titulo VARCHAR(255) NOT NULL,
    titulo_original VARCHAR(255) NULL,
    classificacao_etaria classificacao_etaria NOT NULL,
    descricao VARCHAR(255) NULL,
    edicao VARCHAR(255) NULL,
    idioma VARCHAR(100) NULL,
    preco_base numeric(10,2) NOT NULL
);

CREATE INDEX idx_livro_isbn ON livro(isbn);

CREATE TABLE desconto (
    id SERIAL PRIMARY KEY,
    isbn VARCHAR(255) NOT NULL REFERENCES livro(isbn),
    porcentagem INT CHECK (porcentagem between 0 and 100) NOT NULL,
    data_inicio DATE NOT NULL,
    data_fim DATE NOT NULL CHECK (data_fim > data_inicio),
    evento_atrelado VARCHAR(255) NULL
);

CREATE TYPE genero AS ENUM ('m','f','nb');

CREATE TABLE autor (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    pseudonimo VARCHAR(255) NULL,
    email VARCHAR(255) NOT NULL,
    genero genero NULL,
    data_nascimento DATE NULL,
    data_falecimento DATE NULL CHECK (data_nascimento IS NULL or data_falecimento IS NULL or data_falecimento > data_nascimento),
    nacionalidade VARCHAR(100) NULL
);

CREATE TABLE livro_autor (
    livro_isbn VARCHAR(13) NOT NULL REFERENCES livro(isbn),
    autor_id INT NOT NULL REFERENCES autor(id),
    PRIMARY KEY (livro_isbn, autor_id)
);

CREATE TABLE categoria (
    nome VARCHAR(255) PRIMARY KEY
);

CREATE TABLE livro_categoria (
    livro_isbn VARCHAR(13) NOT NULL REFERENCES livro(isbn),
    categoria_nome VARCHAR(255) NOT NULL REFERENCES categoria(nome),
    PRIMARY KEY (livro_isbn, categoria_nome)
);

CREATE TABLE endereco (
    id SERIAL PRIMARY KEY,
    cep VARCHAR(255) NOT NULL,
    uf VARCHAR(2) NOT NULL,
    localidade VARCHAR(255) NOT NULL,
    logradouro VARCHAR(255) NOT NULL,
    numero INT NOT NULL,
    complemento VARCHAR(255) NULL
);

CREATE TABLE cliente (
    id SERIAL PRIMARY KEY,
    categoria_favorita VARCHAR(255) NULL REFERENCES categoria(nome),
    endereco_id INT REFERENCES endereco(id) NOT NULL,
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    senha VARCHAR(255) NOT NULL,
    data_cadastro DATE NOT NULL,
    data_nascimento DATE NOT NULL
);

CREATE TABLE admin (
    nome VARCHAR(255) NOT NULL,
    senha VARCHAR(255) NOT NULL
);

-- dados de exemplo

INSERT INTO admin VALUES
('ludihan', 'sapomaluco');

INSERT INTO categoria VALUES
('Ficção Científica'),
('Romance'),
('Suspense'),
('Fantasia'),
('Aventura');


INSERT INTO editora (nome, email, data_fundacao, royalties) VALUES
('Editora A', 'editora_a@email.com', '2000-01-01', 15.5),
('Editora B', 'editora_b@email.com', '1995-05-15', 10.0),
('Editora C', 'editora_c@email.com', NULL, 12.3),
('Editora D', 'editora_d@email.com', '2010-12-30', 8.8),
('Editora E', 'editora_e@email.com', '1980-06-20', 20.0);


INSERT INTO livro (isbn, editora_id, titulo, titulo_original, classificacao_etaria, descricao, edicao, idioma, preco_base) VALUES
('1234567890123', 1, 'Livro A', 'Original A', '16', 'Descrição A', '1ª Edição', 'pt', 29.99),
('2345678901234', 2, 'Livro B', NULL, '14', 'Descrição B', '2ª Edição', 'en', 19.99),
('3456789012345', 3, 'Livro C', 'Original C', '10', 'Descrição C', '3ª Edição', 'es', 39.99),
('4567890123456', 1, 'Livro D', 'Original D', '18', 'Descrição D', '4ª Edição', 'pt', 15.99),
('5678901234567', 4, 'Livro E', NULL, '12', 'Descrição E', '5ª Edição', 'fr', 24.99);


INSERT INTO desconto (isbn, porcentagem, data_inicio, data_fim, evento_atrelado) VALUES
('1234567890123', 10, '2023-01-01', '2023-02-01', 'Evento A'),
('2345678901234', 15, '2023-03-01', '2023-04-01', NULL),
('3456789012345', 5, '2023-05-01', '2023-06-01', 'Evento B'),
('4567890123456', 20, '2023-07-01', '2023-08-01', NULL),
('5678901234567', 12, '2023-09-01', '2023-10-01', 'Evento C');


INSERT INTO autor (nome, pseudonimo, email, genero, data_nascimento, data_falecimento, nacionalidade) VALUES
('Autor A', 'Pseudo A', 'autor_a@email.com', 'm', '1980-01-15', NULL, 'Brasileiro'),
('Autor B', NULL, 'autor_b@email.com', 'f', '1995-08-20', '2022-05-10', 'Americano'),
('Autor C', 'Pseudo C', 'autor_c@email.com', 'nb', '1970-04-05', NULL, 'Francês'),
('Autor D', NULL, 'autor_d@email.com', 'm', '1988-11-30', NULL, 'Alemão'),
('Autor E', 'Pseudo E', 'autor_e@email.com', 'f', '2000-06-10', '2021-12-25', 'Mexicano');

INSERT INTO endereco (id, cep, uf, localidade, logradouro, numero, complemento) VALUES
(1, '12345-678', 'SP', 'São Paulo', 'Rua A', 123, 'Apto 101'),
(2, '54321-876', 'RJ', 'Rio de Janeiro', 'Avenida B', 456, NULL),
(3, '98765-432', 'MG', 'Belo Horizonte', 'Rua C', 789, 'Casa 2'),
(4, '65432-109', 'RS', 'Porto Alegre', 'Avenida D', 1011, 'Bloco 3'),
(5, '21098-765', 'BA', 'Salvador', 'Rua E', 1213, NULL);


INSERT INTO cliente (categoria_favorita, endereco_id, nome, email, senha, data_cadastro, data_nascimento) VALUES
('Ficção Científica', 1, 'Cliente A', 'cliente_a@email.com', 'senhaA123', '2022-03-10', '1990-02-12'),
('Romance', 2, 'Cliente B', 'cliente_b@email.com', 'senhaB456', '2021-06-05', '1985-07-18'),
('Suspense', 3, 'Cliente C', 'cliente_c@email.com', 'senhaC789', '2023-01-20', '2000-11-22'),
('Fantasia', 4, 'Cliente D', 'cliente_d@email.com', 'senhaD012', '2020-09-15', '1978-09-05'),
('Aventura', 5, 'Cliente E', 'cliente_e@email.com', 'senhaE345', '2023-05-28', '1998-04-30');
