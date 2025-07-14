-- name: AddCustomer :execresult
INSERT INTO customers (
    username,
    password_hash,
    first_name,
    last_name,
    email,
    birth_day,
    phone_number,
    address
) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetCustomerByID :one
SELECT * FROM customers
WHERE customer_id = ?
LIMIT 1;

-- name: GetCustomerByUsername :one
SELECT * FROM customers
WHERE username = ?
LIMIT 1;
