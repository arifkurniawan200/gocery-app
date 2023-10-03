-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_carts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    product_id INT REFERENCES products(id),
    quantity INT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_carts;
-- +goose StatementEnd
