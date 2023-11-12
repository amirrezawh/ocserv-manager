-- +goose Up

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username text,
    rx_tx_byte bigserial,
    rx_tx text,
    "limit" bigserial,
    active boolean
    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
