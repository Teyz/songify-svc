-- +goose Up

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE rounds (
    id                      VARCHAR(32)     PRIMARY KEY NOT NULL,
    user_id                 VARCHAR(32)     NOT NULL,
    game_id                 VARCHAR(32)     NOT NULL,
    hint                    INTEGER         NOT NULL DEFAULT 0,
    status                  VARCHAR(32)     NOT NULL DEFAULT 'started',
    has_won                 BOOLEAN         NOT NULL DEFAULT FALSE,
    updated_at              TIMESTAMP(6)    NOT NULL DEFAULT NOW(),
    created_at              TIMESTAMP(6)    NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX rounds_user_id_game_id_idx ON rounds (user_id, game_id);
CREATE TRIGGER set_updated_at BEFORE UPDATE ON rounds FOR EACH ROW EXECUTE FUNCTION set_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE rounds;
-- +goose StatementEnd
