-- +goose Up
create table feeds (
    id VARCHAR(50) PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(2000) UNIQUE NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    FOREIGN KEY (user_id) 
    REFERENCES users (id) 
    ON DELETE CASCADE
);


-- +goose Down
drop table feeds;