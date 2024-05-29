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
CREATE TABLE songs (
    id                      VARCHAR(32)     PRIMARY KEY NOT NULL,
    title                   VARCHAR(255)    NOT NULL,
    artist                  VARCHAR(255)    NOT NULL,
    artist_image_url        TEXT            NOT NULL,
    lyrics                  TEXT            NOT NULL,
    image_url               TEXT            NOT NULL,
    released_year           INTEGER         NOT NULL,
    musical_style           VARCHAR(255)    NOT NULL,
    has_been_daily_used     BOOLEAN         NOT NULL DEFAULT FALSE,
    created_at              TIMESTAMP(6)    NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMP(6)    NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX songs_title_artist_idx ON songs (title, artist);

CREATE TRIGGER set_updated_at BEFORE UPDATE ON songs FOR EACH ROW EXECUTE FUNCTION set_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX songs_title_artist_idx;
DROP TABLE songs;
-- +goose StatementEnd
