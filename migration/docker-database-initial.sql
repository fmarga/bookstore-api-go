create table books(
  id serial primary key,
  name varchar,
  author varchar,
  publication_year int,
  pages int,
  publisher varchar
);

INSERT INTO books(name, author, publication_year, pages, publisher) VALUES(
  'Cem Anos de Solidão', 'Gabriel Garcia Marques', 1967, 419, 'Editora Record'
);