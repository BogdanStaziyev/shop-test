CREATE TABLE IF NOT EXISTS customers(
    id              serial PRIMARY KEY,
    email           varchar(50) NOT NULL,
    password        varchar(100) NOT NULL,
    "name"          varchar(50) NOT NULL,
    phone           varchar(30) NOT NULL,
    description     varchar(250) NOT NULL,
    created_date    timestamp NOT NULL,
    updated_date    timestamp NOT NULL,
    deleted_date    timestamp NULL
);