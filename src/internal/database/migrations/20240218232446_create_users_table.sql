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
    artist_image_url        VARCHAR(255)    NOT NULL,
    lyrics                  TEXT            NOT NULL,
    image_url               VARCHAR(255)    NOT NULL,
    released_year           INTEGER         NOT NULL,
    musical_style           VARCHAR(255)    NOT NULL,
    created_at              TIMESTAMP(6)    NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMP(6)    NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE songs;
-- +goose StatementEnd
