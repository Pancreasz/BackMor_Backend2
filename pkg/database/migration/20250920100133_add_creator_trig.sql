-- +goose Up
-- +goose StatementBegin
-- 1. Create the trigger function
CREATE OR REPLACE FUNCTION add_creator_to_activity_members()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO activity_members (activity_id, user_id, role)
    VALUES (NEW.id, NEW.creator_id, 'creator');
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 2. Create the trigger
CREATE TRIGGER trg_add_creator_to_activity_members
AFTER INSERT ON activities
FOR EACH ROW
EXECUTE FUNCTION add_creator_to_activity_members();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS trg_add_creator_to_activity_members ON activities;
DROP FUNCTION IF EXISTS add_creator_to_activity_members();
-- +goose StatementEnd
