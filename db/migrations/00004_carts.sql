-- +goose Up
-- +goose StatementBegin
CREATE TABLE carts (
    cart_id INT PRIMARY KEY AUTO_INCREMENT,
    customer_id INT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers (
        customer_id
    ) ON DELETE CASCADE
);

CREATE TABLE cart_items (
    cart_id INT,
    product_id INT,
    quantity INT NOT NULL CHECK (quantity >= 1),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (cart_id, product_id),
    FOREIGN KEY (cart_id) REFERENCES carts (cart_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (product_id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart_items;
DROP TABLE carts;
-- +goose StatementEnd
