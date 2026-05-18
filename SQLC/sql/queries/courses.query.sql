-- name: ListCourses :many
SELECT c.*, ca.name as category_name FROM courses c
JOIN categories ca ON ca.id = c.category_id
;

-- name: GetCourse :one
SELECT * FROM courses
WHERE ID = ?;

-- name: CreateCourse :exec
INSERT INTO courses (id, category_id, name, description, price)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateCourse :exec
UPDATE courses SET name = ?, description = ?, price = ?
WHERE ID = ?;

-- name: DeleteCourse :exec
DELETE FROM courses WHERE ID = ?;