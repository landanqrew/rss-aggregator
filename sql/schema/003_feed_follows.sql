-- +goose Up
create table feed_follows (
    id VARCHAR(50) PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    feed_id VARCHAR(50) NOT NULL,
    user_id VARCHAR(50) NOT NULL,
    FOREIGN KEY (feed_id) 
    REFERENCES feeds (id) 
    ON DELETE CASCADE,
    FOREIGN KEY (user_id) 
    REFERENCES users (id) 
    ON DELETE CASCADE,
    CONSTRAINT unique_feed_follows UNIQUE (feed_id, user_id)
);


-- +goose Down
drop table feed_follows;