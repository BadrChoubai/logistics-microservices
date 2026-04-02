CREATE TABLE security_reading (
    id uuid PRIMARY KEY,
    sensor_id uuid NOT NULL,
    CONSTRAINT fk_security_reading_sensor FOREIGN KEY (sensor_id) REFERENCES sensor(id),
    door_open boolean NOT NULL DEFAULT false,
    recorded_at timestamptz NOT NULL,

    created_at timestamptz DEFAULT now()
);
