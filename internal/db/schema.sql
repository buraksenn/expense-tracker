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
    installment integer,
    installment_end_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- index for expenses table
CREATE INDEX IF NOT EXISTS expenses_user_id_idx ON expenses(user_id);
CREATE INDEX IF NOT EXISTS expenses_created_at_idx ON expenses(created_at);
CREATE INDEX IF NOT EXISTS expenses_installment_end_date_idx ON expenses(installment_end_date);
CREATE INDEX IF NOT EXISTS expenses_type_idx ON expenses(type);