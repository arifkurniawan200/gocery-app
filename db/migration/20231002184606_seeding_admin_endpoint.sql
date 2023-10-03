-- +goose Up
-- +goose StatementBegin
INSERT INTO users(username,password_hash,is_admin) VALUES ('admin','$2a$10$7ldEb4i4LMQA2LpsuHidneJIleJxX8iXQddplmfXNdg7KfRtjH/n6',true);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
