CREATE TABLE IF NOT EXISTS groups
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS songs
(
    id           SERIAL PRIMARY KEY,
    name         TEXT NOT NULL,
    created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    release_date TIMESTAMP,
    group_id     INT REFERENCES groups (id) ON DELETE CASCADE
);
