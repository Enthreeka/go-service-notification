CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- CREATE TABLE IF NOT EXISTS client_properties(
--     id uuid DEFAULT uuid_generate_v4(),
--     phone_number char(11) unique not null,
--     operator_code char(3) not null,
--     primary key (id)
-- );

CREATE TABLE IF NOT EXISTS notification(
    id uuid DEFAULT uuid_generate_v4(),
    operator_code char(3) not null,
    tag varchar(20),
    created_at timestamp not null,
    message text not null,
    expires_at timestamp not null,
    primary key (id)
);

CREATE INDEX idx_created_at ON notification(created_at);

CREATE TABLE IF NOT EXISTS client(
    id uuid DEFAULT uuid_generate_v4(),
    phone_number char(11) unique not null,
    operator_code char(3) not null,
    tag varchar(20),
    time_zone timestamptz not null,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS message(
    id uuid DEFAULT uuid_generate_v4(),
    notification_id uuid not null,
    client_id uuid not null,
    created_at timestamp not null,
    status varchar(20) not null,
    primary key (id),
    foreign key (notification_id)
        references notification (id),
    foreign key (client_id)
        references client (id)
);

SELECT
    id,
    message,
    created_at,
    expires_at,
    json_agg(json_build_object('tag', tag, 'operator_code', operator_code)) AS client_property
FROM
    notification
WHERE
        created_at = '2023-10-01 20:17:00.000000'
GROUP BY
    id, message, created_at, expires_at;