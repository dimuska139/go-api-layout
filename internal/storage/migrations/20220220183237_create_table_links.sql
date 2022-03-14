-- +goose Up
-- +goose StatementBegin
CREATE TABLE links (
   id BIGSERIAL PRIMARY KEY,
   full_url text NOT NULL,
   code text UNIQUE,
   created_at timestamptz NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE links;
-- +goose StatementEnd
