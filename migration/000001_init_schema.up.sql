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
    primary key (id),
    foreign key (notification_id)
        references notification (id),
    foreign key (client_id)
        references client (id)
);

SELECT notification.id,notification.created_at,notification.message,notification.expires_at,
       client.id,client.phone_number
FROM notification
         JOIN client ON notification.id_client_properties = client.id_client_properties
WHERE notification.created_at = $1;