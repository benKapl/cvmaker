-- +goose Up
CREATE TABLE offers(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    label TEXT UNIQUE NOT NULL,
    organization TEXT NOT NULL,
    organization_description TEXT,
    missions TEXT NOT NULL,
    stack TEXT,
    expected_profile TEXT NOT NULL,
    miscellaneous TEXT,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE offers;