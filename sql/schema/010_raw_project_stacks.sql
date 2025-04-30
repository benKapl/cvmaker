-- +goose Up
CREATE TABLE raw_project_stacks(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    project_id UUID NOT NULL REFERENCES raw_projects(id) ON DELETE CASCADE,
    stack_id UUID NOT NULL REFERENCES raw_stacks(id) ON DELETE CASCADE,
    unique(project_id, stack_id)
);

-- +goose Down
DROP TABLE raw_project_stacks;