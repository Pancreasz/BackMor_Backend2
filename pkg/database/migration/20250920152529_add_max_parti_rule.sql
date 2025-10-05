-- +goose Up
-- +goose StatementBegin
CREATE OR REPLACE FUNCTION enforce_max_participants()
RETURNS TRIGGER AS $$
DECLARE
    current_count INT;
    max_allowed   INT;
BEGIN
    -- Get current member count and max allowed for this activity
    SELECT COUNT(am.user_id), a.max_participants
    INTO current_count, max_allowed
    FROM activities a
    LEFT JOIN activity_members am ON am.activity_id = a.id
    WHERE a.id = NEW.activity_id
    GROUP BY a.max_participants;

    -- If the activity is full, block the insert
    IF current_count >= max_allowed THEN
        RAISE EXCEPTION 'Activity is full (max % participants)', max_allowed
            USING ERRCODE = 'check_violation';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_enforce_max_participants
BEFORE INSERT ON activity_members
FOR EACH ROW
EXECUTE FUNCTION enforce_max_participants();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS trg_enforce_max_participants ON activity_members;
DROP FUNCTION IF EXISTS enforce_max_participants();
-- +goose StatementEnd
