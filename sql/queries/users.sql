-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, email, password)
VALUES (
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;