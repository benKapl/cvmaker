-- name: CreateRawStack :one
INSERT INTO raw_stacks(id, created_at, updated_at, label, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: GetRawStackByLabel :one
SELECT * FROM raw_stacks
WHERE label = $1 AND user_id = $2;