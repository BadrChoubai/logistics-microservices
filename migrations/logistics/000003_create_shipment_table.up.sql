CREATE TYPE shipment_status AS ENUM ('warehouse_location', 'in_transit', 'final_destination');

CREATE TABLE shipment (
    id uuid primary key,
    order_date timestamptz NOT NULL,
    status shipment_status NOT NULL,
    origin VARCHAR(255) NOT NULL,
    destination VARCHAR(255) NOT NULL,

    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now()
);

-- @trigger set_shipment_updated_at_trigger
-- @desc updates value of `updated_at` before UPDATE on TABLE(shipment)
CREATE TRIGGER set_shipment_updated_at
    BEFORE UPDATE ON shipment
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
