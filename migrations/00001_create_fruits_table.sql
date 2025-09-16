-- +goose Up
-- +goose StatementBegin
CREATE TABLE fruits (
                                 id SERIAL PRIMARY KEY,
                                 name TEXT not null,
                                 colour TEXT not null,
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE fruits;
-- +goose StatementEnd
