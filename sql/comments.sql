CREATE TABLE comments
(
	id int not null auto_increment,
    email varchar(100) not null,
    comment text,
    primary key (id)
) engine = innodb;

ALTER TABLE comments 
ADD unique key email (email);

DROP TABLE comments;

SELECT * FROM comments;
SELECT * FROM comments where email = "repository@test.com";  
SELECT * FROM comments WHERE id > 110020;

DELETE FROM comments WHERE email = "repository@test.com";

INSERT INTO comments (email, comment) VALUES ("repository@test.com", "a");