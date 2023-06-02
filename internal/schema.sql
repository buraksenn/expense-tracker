CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email text NOT NULL UNIQUE,
    chat_id text NOT NULL UNIQUE,
    spreadsheet_id text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS expenses (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL REFERENCES users(id),
    name text NOT NULL,
    type text NOT NULL,
    price real NOT NULL,
    tax_percentage real NOT NULL,  
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);