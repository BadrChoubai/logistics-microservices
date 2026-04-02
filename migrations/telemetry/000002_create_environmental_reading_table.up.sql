CREATE TABLE environmental_reading (
    id uuid PRIMARY KEY,
    sensor_id uuid NOT NULL,
    CONSTRAINT fk_environmental_reading_sensor FOREIGN KEY (sensor_id) REFERENCES sensor(id),
    temperature_celsius int NOT NULL,
    recorded_at timestamptz NOT NULL,

    created_at timestamptz DEFAULT now()
);

