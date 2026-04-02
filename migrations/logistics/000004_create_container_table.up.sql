CREATE TABLE container (
                           id uuid PRIMARY KEY,
                           shipment_id uuid NOT NULL,
                           CONSTRAINT fk_container_shipment FOREIGN KEY (shipment_id) REFERENCES shipment(id),
                           container_number VARCHAR(11) NOT NULL,

                           created_at timestamptz DEFAULT now(),
                           updated_at timestamptz DEFAULT now()
);

-- @trigger set_container_updated_at_trigger
-- @desc updates value of `updated_at` before UPDATE on TABLE(container)
CREATE TRIGGER set_container_updated_at
    BEFORE UPDATE ON container
    FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
