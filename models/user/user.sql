CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(50) NOT NULL,
                       username VARCHAR(250) NOT NULL UNIQUE,
                       password VARCHAR(50) NOT NULL,
                       email VARCHAR(50) NOT NULL UNIQUE,
                       avatar VARCHAR(200) DEFAULT 'test.png',
                       signature VARCHAR(100) DEFAULT NULL
);