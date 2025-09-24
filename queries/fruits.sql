-- name: ListFruits :many
SELECT * FROM fruits
ORDER BY name;

-- name: InsertFruit :one
INSERT INTO fruits (name, colour) VALUES (?, ?)
RETURNING *;

-- name: DeleteFruit :exec
DELETE from fruits WHERE id = ?;



