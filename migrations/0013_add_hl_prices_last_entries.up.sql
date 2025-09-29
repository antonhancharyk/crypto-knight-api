ALTER TABLE last_entries ADD COLUMN high_prices NUMERIC[] DEFAULT '{}';
ALTER TABLE last_entries ADD COLUMN low_prices NUMERIC[] DEFAULT '{}';
