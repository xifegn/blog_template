CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(50) NOT NULL UNIQUE,
                       password VARCHAR(50) NOT NULL,
                       email VARCHAR(50) NOT NULL UNIQUE,
                       avatar VARCHAR(200),
                       signature VARCHAR(100) DEFAULT NULL
);