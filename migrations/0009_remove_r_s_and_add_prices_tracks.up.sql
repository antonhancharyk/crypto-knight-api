ALTER TABLE tracks DROP COLUMN resistance_price_1;
ALTER TABLE tracks DROP COLUMN support_price_1;
ALTER TABLE tracks DROP COLUMN resistance_price_2;
ALTER TABLE tracks DROP COLUMN support_price_2;
ALTER TABLE tracks DROP COLUMN resistance_price_3;
ALTER TABLE tracks DROP COLUMN support_price_3;
ALTER TABLE tracks ADD COLUMN high_prices NUMERIC[] DEFAULT '{}';
ALTER TABLE tracks ADD COLUMN low_prices NUMERIC[] DEFAULT '{}';
