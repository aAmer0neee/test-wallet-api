CREATE TABLE IF NOT EXISTS wallets (
    id UUID PRIMARY KEY,
    balance DECIMAL NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_wallets_balance ON wallets(balance);