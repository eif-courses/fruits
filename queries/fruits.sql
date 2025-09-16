-- name: ListFruits :many
SELECT * FROM fruits
ORDER BY name;

-- name: InsertFruit :one
INSERT INTO fruits (name, colour) VALUES ($1, $2)
RETURNING *;