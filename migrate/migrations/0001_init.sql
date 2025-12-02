CREATE table telemetry_device (
    id serial primary key,
    uuid varchar(36) not null unique,
    name varchar(255) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

CREATE TABLE telemetry_readings (
    id serial primary key,
    device_id integer not null references telemetry_device(id) on delete cascade,
    type varchar(50) not null,
    value double precision not null,
    recorded_at timestamp not null default current_timestamp
);