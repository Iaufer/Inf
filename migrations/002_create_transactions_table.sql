CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,                          -- id транзакции
    from_address TEXT NOT NULL,                     -- адрес отправителя
    to_address TEXT NOT NULL,                       -- адрес получателя
    amount NUMERIC(10, 2) NOT NULL,                 -- сумма перевода
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  -- время создания транзакции
);
