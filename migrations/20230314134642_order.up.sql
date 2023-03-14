CREATE TABLE IF NOT EXISTS orders(
    id              serial    PRIMARY KEY,
    created_date    timestamp NOT NULL,
    updated_date    timestamp NOT NULL,
    deleted_date    timestamp NULL
);