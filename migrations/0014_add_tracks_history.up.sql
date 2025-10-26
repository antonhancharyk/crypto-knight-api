CREATE TABLE IF NOT EXISTS tracks_history (
    id SERIAL PRIMARY KEY,
    symbol VARCHAR(255) NOT NULL,
    high_price NUMERIC NOT NULL,
    low_price NUMERIC NOT NULL,
    high_prices NUMERIC[] NOT NULL DEFAULT '{}',
    low_prices NUMERIC[] NOT NULL DEFAULT '{}',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
