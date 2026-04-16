CREATE TABLE shipment_item (
                           id uuid PRIMARY KEY,
                           shipment_id uuid NOT NULL,
                           CONSTRAINT fk_shipment_item_shipment FOREIGN KEY (shipment_id) REFERENCES shipment(id),
                           created_at timestamptz DEFAULT now(),
                           updated_at timestamptz DEFAULT now()
);

-- @trigger set_shipment_item_updated_at_trigger
-- @desc updates value of `updated_at` before UPDATE on TABLE(shipment_item)
CREATE TRIGGER set_shipment_item_updated_at
    BEFORE UPDATE ON shipment_item
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
