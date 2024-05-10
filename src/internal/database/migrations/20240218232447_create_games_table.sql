-- +goose Up

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE games (
    id                      VARCHAR(32)     PRIMARY KEY NOT NULL,
    song_id                 VARCHAR(32)     NOT NULL,
    created_at              TIMESTAMP(6)    NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE games;
-- +goose StatementEnd
