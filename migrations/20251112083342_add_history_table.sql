-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS item_history (
    history_id SERIAL PRIMARY KEY,
    item_id INT NOT NULL,
    action VARCHAR(10) NOT NULL,
    changed_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    changed_by INT,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    price FLOAT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS item_history;
-- +goose StatementEnd
