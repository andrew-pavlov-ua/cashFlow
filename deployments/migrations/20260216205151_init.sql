-- +goose Up
-- +goose StatementBegin

-- Accounts table (bank accounts)
CREATE TABLE IF NOT EXISTS clients (
    external_id VARCHAR(255) NOT NULL PRIMARY KEY, -- unique identifier from the bank's API
    source VARCHAR(50) NOT NULL DEFAULT 'monobank', -- 'monobank', 'privatbank'
    tg_chat_id BIGINT NOT NULL, -- Telegram chat ID for notifications
    mono_user_key VARCHAR(255), -- for monobank API access (if applicable)
    
    name VARCHAR(255),
    balance BIGINT NOT NULL DEFAULT 0, -- in minor units (kopiykas/cents)
    
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    
    CONSTRAINT unique_external_account UNIQUE(source, external_id)
);

-- acounts (cards) table (linked to clients)
CREATE TABLE IF NOT EXISTS accounts (
    id ID PRIMARY KEY, -- external account id from the bank's API
    account_id VARCHAR(255) NOT NULL REFERENCES clients(external_id) ON DELETE CASCADE,
    source VARCHAR(50) NOT NULL DEFAULT 'monobank', -- 'monobank', 'privatbank'
    name VARCHAR(255), -- 'black, white etc'
    last4 VARCHAR(4), -- last 4 digits of the card number for display

    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
)

-- jars table (also linked to clients)
CREATE TABLE IF NOT EXISTS jars (
    id ID PRIMARY KEY, -- mono jar id from API
    account_id VARCHAR(255) NOT NULL REFERENCES clients(external_id) ON DELETE CASCADE,
    source VARCHAR(50) NOT NULL DEFAULT 'monobank', -- 'monobank', 'privatbank'
    title VARCHAR(255),
    description TEXT,
    balance BIGINT NOT NULL DEFAULT 0, -- in minor units (kopiykas/cents)
    goal BIGINT, -- in minor units (kopiykas/cents), optional

    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
)

-- Transactions table
CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id UUID NOT NULL REFERENCES accounts(external_id) ON DELETE CASCADE,
    source VARCHAR(50) NOT NULL, -- 'monobank', 'privatbank'
    
    amount BIGINT NOT NULL, -- in minor units (kopiykas/cents)
    currency_code INTEGER NOT NULL,
    
    description TEXT,
    mcc INTEGER, -- Merchant Category Code
    category VARCHAR(100),
    
    transaction_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    
    CONSTRAINT unique_external_transaction UNIQUE(source, external_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS accounts;
+goose StatementEnd
