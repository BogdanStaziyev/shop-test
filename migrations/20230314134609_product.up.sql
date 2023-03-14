CREATE TABLE IF NOT EXISTS products(
    id              serial PRIMARY KEY,
    "name"          varchar(50) NOT NULL,
    description     varchar(250) NOT NULL,
    price           decimal NOT NULL,
    created_date    timestamp NOT NULL,
    updated_date    timestamp NOT NULL,
    deleted_date    timestamp NULL
);