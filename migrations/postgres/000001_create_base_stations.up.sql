CREATE TYPE base_station_status AS ENUM ('ACTIVE', 'INACTIVE');

CREATE TYPE coordinate_system AS ENUM (
    'COORD_SYS_UNSPECIFIED',
    'WGS84',
    'MSK',
    'EPSG_3857',
    'ETRS89',
    'NAD83',
    'PULKOVO_1942',
    'GSK_2011',
    'UTM_ZONE_33N',
    'LOCAL_GRID'
);

CREATE TYPE constellation AS ENUM ('GPS', 'GLONASS', 'GALILEO', 'BEIDOU', 'QZSS', 'SBAS');

CREATE TABLE base_stations
(
    id                   UUID PRIMARY KEY,
    name                 TEXT                NOT NULL,
    position             geometry(PointZ, 4326) NOT NULL,
    coordinate_system    coordinate_system   NOT NULL DEFAULT 'COORD_SYS_UNSPECIFIED',
    coordinate_subsystem TEXT,
    antenna_type         TEXT                NOT NULL,
    antenna_height       DOUBLE PRECISION    NOT NULL,
    receiver_type        TEXT                NOT NULL,
    receiver_serial      TEXT                NOT NULL,
    firmware_version     TEXT                NOT NULL,
    constellations       constellation[] NOT NULL,
    status               base_station_status NOT NULL DEFAULT 'INACTIVE',
    last_calibrated_at   TIMESTAMP,
    installed_at         TIMESTAMP,
    updated_at           TIMESTAMP DEFAULT now()
);