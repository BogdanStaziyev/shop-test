CREATE TABLE IF NOT EXISTS orders_products(
    order_id INTEGER REFERENCES orders(id),
    product_id INTEGER REFERENCES products(id),
    PRIMARY KEY (order_id, product_id)
)