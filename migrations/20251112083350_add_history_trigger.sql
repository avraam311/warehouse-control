-- +goose Up
-- +goose StatementBegin

CREATE OR REPLACE FUNCTION log_item_changes() RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN
        INSERT INTO item_history (item_id, action, changed_at, changed_by, name, description, price)
        VALUES (NEW.id, 'INSERT', now(), current_setting('myapp.current_user_id', true)::int, NEW.name, NEW.description, NEW.price);
        RETURN NEW;
    ELSIF TG_OP = 'UPDATE' THEN
        INSERT INTO item_history (item_id, action, changed_at, changed_by, name, description, price)
        VALUES (OLD.id, 'UPDATE', now(), current_setting('myapp.current_user_id', true)::int, NEW.name, NEW.description, NEW.price);
        RETURN NEW;
    ELSIF TG_OP = 'DELETE' THEN
        INSERT INTO item_history (item_id, action, changed_at, changed_by, name, description, price)
        VALUES (OLD.id, 'DELETE', now(), current_setting('myapp.current_user_id', true)::int, OLD.name, OLD.description, OLD.price);
        RETURN OLD;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_log_item_changes
AFTER INSERT OR UPDATE OR DELETE ON item
FOR EACH ROW
EXECUTE FUNCTION log_item_changes();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TRIGGER IF EXISTS trg_log_item_changes ON item;
DROP FUNCTION IF EXISTS log_item_changes();

-- +goose StatementEnd
