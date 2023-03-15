CREATE TABLE IF NOT EXISTS products(
    id              serial PRIMARY KEY,
    "name"          VARCHAR(50) NOT NULL,
    salesman_id     INTEGER NOT NULL REFERENCES salesmans(id),
    description     VARCHAR(250) NOT NULL,
    price           DECIMAL NOT NULL,
    created_date    TIMESTAMP NOT NULL,
    updated_date    TIMESTAMP NOT NULL,
    deleted_date    TIMESTAMP NULL
);