CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    `uuid` CHAR(36) NOT NULL UNIQUE,
    user_id INTEGER NOT NULL,
    device_timestamp TIMESTAMP NOT NULL,
    total_amount INTEGER NOT NULL DEFAULT 0,
    paid_amount INTEGER NOT NULL DEFAULT 0,
    change_amount INTEGER NOT NULL DEFAULT 0,
    payment_method ENUM('none', 'cash','card') NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS transaction_items (
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    `uuid` CHAR(36) NOT NULL UNIQUE,
    transaction_id INTEGER NOT NULL,
    title VARCHAR(64) NOT NULL,
    qty INTEGER NOT NULL DEFAULT 0,
    price INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT FK_transactionID FOREIGN KEY (transaction_id)
    REFERENCES transactions(id)
);

