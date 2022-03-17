CREATE TABLE customer(
	id varchar(100) not null,
    name varchar(100) not null,
    primary key (id)
) engine = innodb;

SELECT * FROM customer;

DELETE FROM customer;

ALTER TABLE customer 
add column email varchar(100),
add column balance int default 0,
add column rating double default 0.0,
add column created_at timestamp default current_timestamp,
add column birth_date date,
add column married boolean default false;

DESCRIBE customer;

INSERT INTO customer(id, name, email, balance, married, rating, birth_date) VALUES
("p1", "fahmi", "fahmi@gmail.com", 1000000, true, 5.0, '2001-09-29'),
("p2", "ega", "ega@gmail.com", 500000, true, 4.5, '2002-09-29'),
("p3", "soni", "soni@gmail.com", 800000, false, 4.0, '2003-09-29');

UPDATE customer 
SET email = null, birth_date = null
WHERE id = 'p1'