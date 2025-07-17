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
WHERE customer_id = ?;

-- name: GetCustomerByUsername :one
SELECT * FROM customers
WHERE username = ?;

-- name: GetCustomerByToken :one
SELECT c.*
FROM customers AS c
INNER JOIN tokens AS t
    ON c.customer_id = t.customer_id
    -- TODO: might replace the expiry comparison with the internal `> NOW()`
WHERE t.token_hash = ? AND t.token_scope = ? AND t.expiry > sqlc.arg('date');
