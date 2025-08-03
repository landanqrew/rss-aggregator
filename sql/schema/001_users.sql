-- +goose Up
create table users (
    id VARCHAR(50) PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL
);

-- +goose Down
drop table users;