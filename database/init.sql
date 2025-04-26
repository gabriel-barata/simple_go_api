CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(50) NOT NULL,
    price NUMERIC(10, 2) NOT NULL
);

INSERT INTO products (product_name, price) VALUES ('tomato', 2.35);
