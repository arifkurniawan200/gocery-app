-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  product_name VARCHAR(255) NOT NULL,
  category_id INT REFERENCES product_categories(id),
  description TEXT,
  qty INT,
  selling_price DECIMAL,
  promo_price DECIMAL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
