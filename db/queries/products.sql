-- name: GetAllProducts :many
SELECT * FROM products_with_category;

-- name: GetLatestProducts :many
SELECT p.*
FROM products_with_category AS p
WHERE (
    SELECT COUNT(*)
    FROM products_with_category AS p2
    WHERE
        p2.category = p.category
        AND (
            p2.created_at > p.created_at
            OR (p2.created_at = p.created_at AND p2.product_id > p.product_id)
        )
) < CAST(sqlc.arg('category_limit') AS SIGNED)
ORDER BY p.category ASC, p.created_at DESC, p.product_id DESC;

-- name: GetProduct :one
SELECT * FROM products_with_category
WHERE product_id = ?;

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
