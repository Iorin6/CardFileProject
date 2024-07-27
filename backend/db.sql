CREATE TABLE books (
	id serial primary key,
	name VARCHAR(100) NOT NULL,
	author VARCHAR(100),
	year DATE,
	genre VARCHAR(100),
	description VARCHAR(1000)
);