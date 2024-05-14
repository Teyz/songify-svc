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
CREATE TABLE guesses (
    id                      VARCHAR(32)     PRIMARY KEY NOT NULL,
    user_id                 VARCHAR(32)     NOT NULL,
    artist                  VARCHAR(255)    NOT NULL,
    title                   VARCHAR(255)    NOT NULL,
    created_at              TIMESTAMP(6)    NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON songs FOR EACH ROW EXECUTE FUNCTION set_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE guesses;
-- +goose StatementEnd
