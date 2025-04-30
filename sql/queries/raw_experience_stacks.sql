-- name: CreateRawExperienceStack :one
WITH inserted_raw_experience_stack as(
    INSERT INTO raw_experience_stacks(id, created_at, updated_at, experience_id, stack_id)
    VALUES (
        gen_random_uuid(),
        NOW(),
        NOW(),
        $1,
        $2
    )
    RETURNING *
)
    SELECT 
        inserted_raw_experience_stack.*,
        raw_experiences.title AS experience_title,
        raw_stacks.label as stack_label
    FROM inserted_raw_experience_stack
    INNER JOIN raw_experiences on raw_experiences.id = raw_experience_stacks.experience_id
    INNER JOIN raw_stacks on raw_stacks.id = raw_experience_stacks.stack_id;