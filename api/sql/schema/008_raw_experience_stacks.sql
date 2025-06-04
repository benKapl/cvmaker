-- +goose Up
CREATE TABLE raw_experience_stacks(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    experience_id UUID NOT NULL REFERENCES raw_experiences(id) ON DELETE CASCADE,
    stack_id UUID NOT NULL REFERENCES raw_stacks(id) ON DELETE CASCADE,
    unique(experience_id, stack_id)
);

-- +goose Down
DROP TABLE raw_experience_stacks;