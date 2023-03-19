CREATE TABLE IF NOT EXISTS orders(
    id              serial    PRIMARY KEY,
    customer_id     INTEGER   NOT NULL REFERENCES customers(id),
    created_date    TIMESTAMP NOT NULL,
    updated_date    TIMESTAMP NOT NULL,
    deleted_date    TIMESTAMP NULL
);