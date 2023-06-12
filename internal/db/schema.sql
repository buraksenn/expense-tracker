CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    telegram_id text NOT NULL UNIQUE,
    spreadsheet_id text,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS users_telegram_id_idx ON users(telegram_id);

CREATE TABLE IF NOT EXISTS expenses (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL REFERENCES users(id),
    description text NOT NULL,
    type text NOT NULL,
    price real NOT NULL,
    tax_percentage integer NOT NULL,  
    installment integer,
    installment_end_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS expenses_user_id_idx ON expenses(user_id);
CREATE INDEX IF NOT EXISTS expenses_created_at_idx ON expenses(created_at);
CREATE INDEX IF NOT EXISTS expenses_installment_end_date_idx ON expenses(installment_end_date);
CREATE INDEX IF NOT EXISTS expenses_type_idx ON expenses(type);