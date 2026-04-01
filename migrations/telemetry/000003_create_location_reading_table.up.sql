CREATE TABLE location_reading (
    id uuid PRIMARY KEY,
    sensor_id uuid NOT NULL,
    CONSTRAINT fk_location_reading_sensor FOREIGN KEY (sensor_id) REFERENCES sensor(id),
    latitude DECIMAL(9,6) NOT NULL,
    longitude DECIMAL(9,6) NOT NULL,
    recorded_at timestamptz NOT NULL,

    created_at timestamptz DEFAULT now()
);
