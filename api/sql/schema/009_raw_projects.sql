-- +goose Up
CREATE TABLE raw_projects(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    label TEXT NOT NULL,
    description TEXT NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP, -- null end_date means ongoing
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    unique(label, user_id)
);

-- +goose Down
DROP TABLE raw_projects;