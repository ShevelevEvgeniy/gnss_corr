CREATE TYPE rover_type AS ENUM ('GNSS', 'RTK', 'STATIC', 'MOBILE');

CREATE TABLE rovers
(
    id                      UUID PRIMARY KEY,
    owner_id                UUID       NOT NULL,
    name                    TEXT       NOT NULL,
    rover_type              rover_type NOT NULL,
    serial_number           TEXT       NOT NULL,
    subscription_expires_at TIMESTAMP,
    is_active               BOOLEAN   DEFAULT true,
    max_distance_km         DOUBLE PRECISION,
    constellations          constellation[] NOT NULL,
    registered_at           TIMESTAMP DEFAULT now(),
    updated_at              TIMESTAMP DEFAULT now()
);
