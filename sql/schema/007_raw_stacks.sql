-- +goose Up
CREATE TABLE raw_stacks(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    label TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    unique(label, user_id)
);

-- +goose Down
DROP TABLE raw_stacks;