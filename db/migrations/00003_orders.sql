-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    order_id INT PRIMARY KEY AUTO_INCREMENT,
    customer_id INT,
    order_status ENUM(
        'pending', 'processing', 'shipped', 'delivered', 'cancelled'
    ) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    ON UPDATE CURRENT_TIMESTAMP,
    order_total DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers (
        customer_id
    ) ON DELETE SET NULL
);

CREATE INDEX idx_orders_placed ON orders (created_at);

CREATE TABLE order_items (
    order_id INT,
    product_id INT,
    quantity INT NOT NULL CHECK (quantity >= 1),
    -- TODO: made the price below populate at creation from the product price
    product_price_at_time DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders (order_id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (product_id) ON DELETE CASCADE
);

-- CREATE TRIGGER after_order_item_insert
-- AFTER INSERT ON order_items
-- FOR EACH ROW
-- UPDATE products
-- SET stock_quantity = stock_quantity - new.quantity
-- WHERE product_id = new.product_description;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_items;
DROP TABLE orders;
-- +goose StatementEnd
