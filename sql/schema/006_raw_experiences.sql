-- +goose Up
CREATE TABLE raw_experiences(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    organization TEXT NOT NULL,
    description TEXT NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP, -- null end_date means ongoing
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    unique(title, organization, user_id)
);

-- +goose Down
DROP TABLE raw_experiences;