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
    title = ?,
    product_description = ?,
    category = ?,
    price = ?,
    image_url = ?,
    stock_quantity = ?
WHERE product_id = ?;
