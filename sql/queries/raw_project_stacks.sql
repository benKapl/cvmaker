-- name: CreateRawprojectStack :one
WITH inserted_raw_project_stack as(
    INSERT INTO raw_project_stacks(id, created_at, updated_at, project_id, stack_id)
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
        inserted_raw_project_stack.*,
        raw_projects.label AS project_label,
        raw_stacks.label AS stack_label
    FROM inserted_raw_project_stack
    INNER JOIN raw_projects ON raw_projects.id = inserted_raw_project_stack.project_id
    INNER JOIN raw_stacks ON raw_stacks.id = inserted_raw_project_stack.stack_id;