-- +goose Up
-- +goose StatementBegin
CREATE TABLE product_categories (
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product_categories;
-- +goose StatementEnd
