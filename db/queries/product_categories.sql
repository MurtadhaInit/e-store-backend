-- name: GetProductCategory :one
SELECT * FROM product_categories
WHERE category_id = ?
LIMIT 1;

-- name: AddProductCategory :execresult
INSERT INTO product_categories (
    category_name, category_description
) VALUES (?, ?);
