-- name: AddToken :execresult
INSERT INTO tokens (
    token_hash,
    customer_id,
    expiry,
    token_scope
) VALUES (?, ?, ?, ?);

-- name: DeleteAllTokensForCustomer :execresult
-- DELETE FROM tokens
-- WHERE token_scope = ? AND customer_id = ?;
