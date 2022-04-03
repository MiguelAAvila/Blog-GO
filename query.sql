--  Members: Miguel Avila, Federico Rosado

DROP TABLE IF EXISTS blogs;

CREATE TABLE blogs (
    id serial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    email text NOT NULL,
    subject text NOT NULL,
    message text NOT NULL,
    created_at date NOT NULL DEFAULT NOW()
);