-- +goose Up
-- +goose StatementBegin
CREATE TABLE product_categories (
    category_id INT PRIMARY KEY AUTO_INCREMENT,
    category_name VARCHAR(50) NOT NULL UNIQUE,
    category_description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE products (
    product_id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL UNIQUE,
    product_description TEXT NOT NULL,
    category INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    stock_quantity INT NOT NULL CHECK (stock_quantity >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category) REFERENCES product_categories (
        category_id
    ) ON DELETE RESTRICT
);

CREATE INDEX idx_product_title ON products (title);

CREATE VIEW products_with_category AS
SELECT
    p.product_id,
    p.title,
    p.product_description,
    pc.category_name AS category,
    p.price,
    p.image_url,
    p.stock_quantity,
    p.created_at,
    p.updated_at
FROM products AS p
INNER JOIN product_categories AS pc ON p.category = pc.category_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW products_with_category;
DROP TABLE products;
DROP TABLE product_categories;
-- +goose StatementEnd
