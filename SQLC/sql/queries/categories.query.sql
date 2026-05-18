-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories
WHERE ID = ?;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description)
VALUES (?, ?, ?);

-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ?
WHERE ID = ?;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE ID = ?;