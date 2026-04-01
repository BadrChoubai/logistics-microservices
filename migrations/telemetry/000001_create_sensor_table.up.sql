CREATE TYPE sensor_type AS ENUM ('environmental', 'location', 'security');

CREATE TABLE sensor (
    id uuid PRIMARY KEY,
    container_id uuid NOT NULL, -- references container(id) in logistics db, not enforced
    type sensor_type NOT NULL,

    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);

-- @function set_updated_at
-- @desc Updates the updated_at column to current timestamp
-- @returns TRIGGER
-- @used_by triggers on all tables where updated_at changes on a frequent basis
CREATE OR REPLACE FUNCTION set_updated_at()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- @trigger set_sensor_updated_at_trigger
-- @desc updates value of `updated_at` before UPDATE on TABLE(sensor)
CREATE TRIGGER set_sensor_updated_at
    BEFORE UPDATE ON sensor
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
