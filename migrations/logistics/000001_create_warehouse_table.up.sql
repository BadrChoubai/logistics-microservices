-- @table warehouse
-- @desc stores warehouse information
CREATE TABLE warehouse (
    id uuid PRIMARY KEY,
    address VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100) NOT NULL,
    zip_code VARCHAR(20) NOT NULL ,

    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);

-- @function set_warehouse_updated_at
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

-- @trigger set_warehouse_updated_at_trigger
-- @desc updates value of `updated_at` before UPDATE on TABLE(warehouse)
CREATE TRIGGER set_warehouse_updated_at
    BEFORE UPDATE ON warehouse
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
