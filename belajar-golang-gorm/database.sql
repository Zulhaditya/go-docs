-- CREATE DATABASE belajar_golang_gorm;
-- USE belajar_golang_gorm;=
-- SHOW TABLES;
CREATE TABLE sample (
    id varchar(100) not null,
    name varchar(100) not null,
    primary key (id)
) ENGINE = InnoDB;

CREATE TABLE users (
    id varchar(100) not null,
    password varchar(100) not null,
    name varchar(100) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id)
) ENGINE = InnoDB;

SELECT
    *
FROM
    users;

ALTER TABLE
    users RENAME COLUMN name to first_name;

ALTER TABLE
    users
ADD
    COLUMN middle_name varchar(100) null
after
    first_name;

ALTER TABLE
    users
ADD
    COLUMN last_name varchar(100) null
after
    middle_name;

CREATE TABLE user_logs (
    id int AUTO_INCREMENT,
    user_id VARCHAR(100) NOT NULL,
    action VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

DELETE FROM user_logs;

ALTER TABLE user_logs
    MODIFY created_at BIGINT NOT NULL;

ALTER TABLE user_logs
    MODIFY updated_at BIGINT NOT NULL;

SELECT * FROM user_logs;

DESC user_logs;

CREATE TABLE todos (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id VARCHAR(100) NOT NULL,
    title VARCHAR(100) NOT NULL,
    description text NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

SELECT * FROM todos;

CREATE TABLE wallets (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id VARCHAR(100) NOT NULL,
    balance BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    Foreign Key (user_id) REFERENCES users (id)
) ENGINE = InnoDB;

DESC wallets;

SELECT * FROM users;

CREATE TABLE addresses (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id VARCHAR(100) NOT NULL,
    address VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    Foreign Key (user_id) REFERENCES users (id)
) ENGINE = InnoDB;

DESC addresses;

SELECT * FROM addresses;