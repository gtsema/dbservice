-- +goose Up
CREATE TABLE users (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   chatId TEXT NOT NULL UNIQUE,
   name TEXT NOT NULL,
   hash TEXT NOT NULL,
   salt TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;