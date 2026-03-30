CREATE TABLE accounts
(
    account_id UUID PRIMARY KEY             DEFAULT gen_random_uuid(),
    user_id    UUID,
    name       VARCHAR(255) NOT NULL,
    balance    BIGINT DEFAULT 0,
    created_at TIMESTAMP           NOT NULL DEFAULT NOW()
);



CREATE INDEX idx_account_name ON accounts (name);