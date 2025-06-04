-- name: CreateRawEducation :one
INSERT INTO raw_educations(id, created_at, updated_at, label, school, description, start_date, end_date, user_id)
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