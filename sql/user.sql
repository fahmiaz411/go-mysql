CREATE TABLE user
(
	username varchar(100) not null,
    password varchar(100) not null,
    primary key (username)
) engine = innodb;

INSERT INTO user (username, password) VALUES
('admin', 'admin');

SELECT * FROM user;
SELECT LAST_INSERT_ID();