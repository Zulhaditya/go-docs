-- CREATE DATABASE belajar_golang_gorm;
-- USE belajar_golang_gorm;=
-- SHOW TABLES;

CREATE TABLE sample
(
	id varchar(100) not null,
    name varchar(100) not null,
    primary key (id)
) ENGINE = InnoDB;