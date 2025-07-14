-- +goose Up
-- +goose StatementBegin
CREATE TABLE customers (
    customer_id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(40) NOT NULL,
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

-- CREATE TABLE tokens (
-- );

CREATE INDEX idx_customer_email ON customers (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE customers;
-- +goose StatementEnd
