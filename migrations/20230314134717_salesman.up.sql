CREATE TABLE IF NOT EXISTS salesmans(
    id              serial PRIMARY KEY,
    email           varchar(50) NOT NULL,
    password        varchar(100) NOT NULL,
    "name"          varchar(50) NOT NULL,
    phone           varchar(30) NOT NULL,
    "role"          varchar(10) NOT NULL,
    created_date    timestamp NOT NULL,
    updated_date    timestamp NOT NULL,
    deleted_date    timestamp NULL
);