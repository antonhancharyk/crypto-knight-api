DROP FUNCTION IF EXISTS update_updated_at_column();
DROP TRIGGER IF EXISTS tracks_update_trigger ON tracks;
DROP TABLE IF EXISTS tracks;