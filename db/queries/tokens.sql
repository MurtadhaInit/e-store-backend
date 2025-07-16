-- name: AddToken :execresult
INSERT INTO tokens (
    token_hash,
    customer_id,
    expiry,
    token_scope
) VALUES (?, ?, ?, ?);

-- name: GetToken :one
SELECT * FROM tokens
WHERE customer_id = ?
LIMIT 1;

-- name: DeleteAllTokensForUser :execresult
DELETE FROM tokens
WHERE token_scope = ? AND customer_id = ?
