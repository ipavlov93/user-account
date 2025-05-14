-- +goose Down
drop table IF NOT EXISTS users;
drop table IF NOT EXISTS user_accounts;
drop table IF NOT EXISTS user_profiles;
