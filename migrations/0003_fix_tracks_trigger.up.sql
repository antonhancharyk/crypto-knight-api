DROP TRIGGER IF EXISTS tracks_update_trigger ON tracks;

CREATE TRIGGER tracks_update_trigger
BEFORE INSERT OR UPDATE ON tracks
FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();