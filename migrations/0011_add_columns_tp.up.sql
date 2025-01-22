ALTER TABLE tracks ADD COLUMN take_profit_high_prices NUMERIC[] DEFAULT '{}';
ALTER TABLE tracks ADD COLUMN take_profit_low_prices NUMERIC[] DEFAULT '{}';
