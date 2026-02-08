-- +goose Up
ALTER TABLE users ADD COLUMN invite_notified BOOLEAN NOT NULL DEFAULT FALSE;

-- +goose Down
ALTER TABLE users DROP COLUMN invite_notified;
