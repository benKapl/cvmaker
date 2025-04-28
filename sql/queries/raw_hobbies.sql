-- name: CreateRawHobby :one
INSERT INTO raw_hobbies(id, created_at, updated_at, label, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;