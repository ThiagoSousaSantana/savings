CREATE TABLE IF NOT EXISTS expense (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    value NUMERIC(10, 2) NOT NULL
);

