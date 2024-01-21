-- CREATE DATABASE belajar_golang_gorm;
-- USE belajar_golang_gorm;=
-- SHOW TABLES;

CREATE TABLE sample
(
	id varchar(100) not null,
    name varchar(100) not null,
    primary key (id)
) ENGINE = InnoDB;

CREATE TABLE users
(
    id varchar(100) not null,
    password varchar(100) not null,
    name varchar(100) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id)
)ENGINE = InnoDB;

SELECT * FROM users;