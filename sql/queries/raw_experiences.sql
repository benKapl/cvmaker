-- name: CreateRawExperience :one
INSERT INTO raw_experiences(id, created_at, updated_at, title, organization, description, start_date, end_date, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;