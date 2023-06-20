CREATE TABLE IF NOT EXISTS manga(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    genre VARCHAR(255),
    volumes INTEGER,
    chapters INTEGER,
    author VARCHAR(255)
);