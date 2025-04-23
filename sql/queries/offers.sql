-- name: CreateOffer :one
INSERT INTO offers(
    id,
    created_at,
    updated_at,
    label,
    organization,
    organization_description,
    missions,
    stack,
    expected_profile,
    miscellaneous,
    user_id
)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8

)
RETURNING *;