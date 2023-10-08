CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS client_properties(
    id uuid DEFAULT uuid_generate_v4(),
    tag char(11) not null,
    operator_code char(3) unique not null,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS notification(
    id uuid DEFAULT uuid_generate_v4(),
    id_client_properties uuid,
    created_at timestamp not null,
    message text not null,
    expires_at timestamp not null,
    primary key (id),
    foreign key (id_client_properties)
        references client_properties (id)
);

CREATE INDEX idx_created_at ON notification(created_at);

CREATE TABLE IF NOT EXISTS client(
    id uuid DEFAULT uuid_generate_v4(),
    id_client_properties uuid,
    phone_number char(11) unique not null,
    time_zone timestamptz not null,
    primary key (id),
    foreign key (id_client_properties)
        references client_properties (id)
);

CREATE TABLE IF NOT EXISTS message(
    id uuid DEFAULT uuid_generate_v4(),
    notification_id uuid not null,
    client_id uuid not null,
    created_at timestamp not null,
    status varchar(20) not null,
    primary key (notification_id,client_id),
    foreign key (notification_id)
        references notification (id),
    foreign key (client_id)
        references client (id)
);


SELECT
    message.status, COUNT(*),
FROM
    message
        INNER JOIN notification  ON message.notification_id = notification.id
GROUP BY
    message.status;

SELECT
    n.id AS notification_id,
    .status AS message_status,
    COUNT(*) AS message_count
FROM
    message
        INNER JOIN
    notification n ON m.notification_id = n.id
GROUP BY
    n.id,
    m.status;

SELECT notification.id, notification.message, message.status, COUNT(*)
    FROM message
    JOIN notification ON notification.id = message.notification_id
GROUP BY
    notification.id, notification.message,message.status;