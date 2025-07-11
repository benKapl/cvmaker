-- +goose Up
CREATE TABLE templates(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    label TEXT NOT NULL,
    description TEXT
);

-- +goose Down
DROP TABLE templates;