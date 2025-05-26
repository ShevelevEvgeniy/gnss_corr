CREATE TABLE allowed_base_access
(
    rover_id UUID,
    base_id  UUID,
    PRIMARY KEY (rover_id, base_id),
    FOREIGN KEY (rover_id) REFERENCES rovers (id),
    FOREIGN KEY (base_id) REFERENCES base_stations (id)
);
