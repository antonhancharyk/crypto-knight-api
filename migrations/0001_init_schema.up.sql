CREATE TABLE IF NOT EXISTS common (
    id SERIAL PRIMARY KEY,
    enabled BOOLEAN NOT NULL
);

INSERT INTO common (enabled) VALUES (true);