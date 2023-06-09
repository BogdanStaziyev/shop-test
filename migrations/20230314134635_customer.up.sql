CREATE TABLE IF NOT EXISTS customers(
    id              serial PRIMARY KEY,
    email           VARCHAR(50) NOT NULL UNIQUE,
    password        VARCHAR(100) NOT NULL,
    "name"          VARCHAR(50) NOT NULL,
    phone           VARCHAR(30) NOT NULL UNIQUE,
    created_date    TIMESTAMP NOT NULL,
    updated_date    TIMESTAMP NOT NULL,
    deleted_date    TIMESTAMP NULL
);