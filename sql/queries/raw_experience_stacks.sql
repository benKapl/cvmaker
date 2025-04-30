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
        raw_stacks.label AS stack_label
    FROM inserted_raw_experience_stack
    INNER JOIN raw_experiences ON raw_experiences.id = inserted_raw_experience_stack.experience_id
    INNER JOIN raw_stacks ON raw_stacks.id = inserted_raw_experience_stack.stack_id;