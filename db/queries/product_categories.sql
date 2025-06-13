-- name: GetProductCategory :one
SELECT * FROM product_categories
WHERE category_id = ?
LIMIT 1;
