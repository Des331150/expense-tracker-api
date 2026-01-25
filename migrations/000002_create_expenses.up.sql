CREATE TABLE IF NOT EXISTS expenses(
    id BIGSERIAL PRIMARY KEY,
    currency TEXT NOT NULL,
    amount BIGINT NOT NULL,
    transaction_type TEXT NOT NULL,
    descriptions TEXT NOT NULL,
    date_of_expense TIMESTAMPTZ,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    user_id BIGINT REFERENCES users (id) ON DELETE CASCADE,
    category_id BIGINT REFERENCES categories (id) ON DELETE SET NULL
);