CREATE TABLE comments
(
	id int not null auto_increment,
    email varchar(100) not null,
    comment text,
    primary key (id)
) engine = innodb;

SELECT * FROM comments;