-- +goose Up
-- +goose StatementBegin
CREATE TABLE customers (
    customer_id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(40) NOT NULL UNIQUE,
    password_hash BINARY(60) NOT NULL,
    first_name VARCHAR(40) NOT NULL,
    last_name VARCHAR(40) NOT NULL,
    email VARCHAR(40) NOT NULL UNIQUE,
    birth_day DATE NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_customer_email ON customers (email);

CREATE TABLE tokens (
    token_hash BINARY(32) PRIMARY KEY,
    customer_id INT NOT NULL,
    expiry TIMESTAMP NOT NULL,
    token_scope TEXT NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers (
        customer_id
    ) ON DELETE CASCADE
);

CREATE INDEX idx_token ON tokens (token_hash);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tokens;
DROP TABLE customers;
-- +goose StatementEnd
