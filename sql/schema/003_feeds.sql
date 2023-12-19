-- +goose Up

CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE --delete all feeds when user is deleted
);

-- +goose Down

DROP TABLE feeds;
