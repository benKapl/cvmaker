-- name: CreateRawProject :one
INSERT INTO raw_projects(id, created_at, updated_at, label, description, start_date, end_date, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;