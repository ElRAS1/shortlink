CREATE TABLE IF NOT EXISTS link (
    id SERIAL PRIMARY KEY,
    oldlink VARCHAR(255) NOT NULL UNIQUE,
    newlink VARCHAR(255) NOT NULL UNIQUE
);