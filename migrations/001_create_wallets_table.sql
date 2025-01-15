CREATE TABLE IF NOT EXISTS wallets (
    id SERIAL PRIMARY KEY,                  -- ID кошелька
    address TEXT UNIQUE NOT NULL,           -- aдрес кошелька
    balance NUMERIC(10, 2) NOT NULL         -- баланс кошелька
);
