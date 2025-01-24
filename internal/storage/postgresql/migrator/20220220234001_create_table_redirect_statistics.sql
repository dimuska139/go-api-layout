-- +goose Up
-- +goose StatementBegin
CREATE TABLE redirect_statistics (
    id BIGSERIAL PRIMARY KEY,
    link_id INT NOT NULL,
    user_agent TEXT,
    ip INET,
    created_at timestamptz NOT NULL,
    CONSTRAINT statistics_links_fkey
        FOREIGN KEY(link_id)
            REFERENCES link(id)
            ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE redirect_statistics;
-- +goose StatementEnd
