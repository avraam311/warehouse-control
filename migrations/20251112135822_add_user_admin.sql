-- +goose Up
-- +goose StatementBegin
INSERT INTO "user" (email, role, hash) VALUES (
    'admin@mail.ru',
    'admin',
    decode('243261243130246462556241714358797732504644306d31556c586575774f4c2f50503071712f6a744c797a3062524d485a4c535139383476494961', 'hex')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM "user" WHERE email = 'admin@mail.ru';
-- +goose StatementEnd
