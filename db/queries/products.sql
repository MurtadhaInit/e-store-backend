-- name: GetAllProducts :many
SELECT * FROM products;

-- name: GetLatestProducts :many
SELECT * FROM products
ORDER BY created_at DESC
LIMIT ?;

-- name: GetProduct :one
SELECT * FROM products
WHERE product_id = ?
LIMIT 1;

-- name: AddProduct :execresult
INSERT INTO products (
    title, product_description, category, price, image_url, stock_quantity
) VALUES (?, ?, ?, ?, ?, ?);

-- name: DeleteProduct :execresult
DELETE FROM products
WHERE product_id = ?;

-- name: EditProduct :execresult
UPDATE products
SET
    title = COALESCE(sqlc.narg('title'), title),
    product_description
    = COALESCE(sqlc.narg('product_description'), product_description),
    category = COALESCE(sqlc.narg('category'), category),
    price = COALESCE(sqlc.narg('price'), price),
    image_url = COALESCE(sqlc.narg('image_url'), image_url),
    stock_quantity = COALESCE(sqlc.narg('stock_quantity'), stock_quantity)
WHERE product_id = sqlc.arg('product_id');
