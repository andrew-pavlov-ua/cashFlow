-- +goose Up
-- +goose StatementBegin

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS clients (
    external_id VARCHAR(255) NOT NULL PRIMARY KEY,
    source VARCHAR(50) NOT NULL DEFAULT 'monobank',
    tg_chat_id BIGINT NOT NULL,
    mono_user_key VARCHAR(255),

    name VARCHAR(255),
    balance BIGINT NOT NULL DEFAULT 0,

    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS accounts (
    id VARCHAR(20) PRIMARY KEY,
    client_external_id VARCHAR(255) NOT NULL REFERENCES clients(external_id) ON DELETE CASCADE,
    source VARCHAR(50) NOT NULL DEFAULT 'monobank',
    name VARCHAR(255),
    last4 VARCHAR(4),

    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS jars (
    id VARCHAR(20) PRIMARY KEY,
    client_external_id VARCHAR(255) NOT NULL REFERENCES clients(external_id) ON DELETE CASCADE,
    source VARCHAR(50) NOT NULL DEFAULT 'monobank',
    title VARCHAR(255),
    description TEXT,
    balance BIGINT NOT NULL DEFAULT 0,
    goal BIGINT,

    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    account_id VARCHAR(20) NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
    source VARCHAR(50) NOT NULL,

    amount BIGINT NOT NULL,
    currency_code INTEGER NOT NULL,

    description TEXT,
    mcc INTEGER,
    category VARCHAR(100),

    transaction_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS jars;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS clients;

-- +goose StatementEnd
