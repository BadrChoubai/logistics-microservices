-- @table inventory
-- @desc stores inventory information
CREATE TABLE inventory (
    id uuid PRIMARY KEY,
    warehouse_id uuid NOT NULL,
    CONSTRAINT fk_inventory_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouse(id),
    name VARCHAR(255) NOT NULL,
    quantity int default 0,

    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);

-- @trigger set_inventory_updated_at_trigger
-- @desc updates value of `updated_at` before UPDATE on TABLE(inventory)
CREATE TRIGGER set_inventory_updated_at
    BEFORE UPDATE ON inventory
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
