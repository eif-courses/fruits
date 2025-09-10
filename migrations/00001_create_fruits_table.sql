-- +goose Up
-- +goose StatementBegin
CREATE TABLE fruits (
                                 id SERIAL PRIMARY KEY,
                                 name VARCHAR(100),
                                 colour VARCHAR(255),
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE fruits;
-- +goose StatementEnd
