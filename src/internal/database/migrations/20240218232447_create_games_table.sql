-- +goose Up

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_updated_at_column() RETURNS TRIGGER AS $$
  BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
  END;
$$ language 'plpgsql';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE games (
    id                      VARCHAR(32)     PRIMARY KEY NOT NULL,
    song_id                 VARCHAR(32)     NOT NULL,
    is_active               BOOLEAN         NOT NULL DEFAULT TRUE,
    created_at              TIMESTAMP(6)    NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE games;
-- +goose StatementEnd
