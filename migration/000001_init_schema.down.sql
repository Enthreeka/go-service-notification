CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS client_properties(
    id uuid DEFAULT uuid_generate_v4(),
    phone_number char(11) not null,
    operator_code char(3) not null,
    primary key (id)
);

CREATE TABLE IF NOT EXISTS notification(
    id uuid DEFAULT uuid_generate_v4(),
    client_properties_id uuid not null,
    created_at timestamp not null,
    message text not null,
    expires_at timestamp not null,
    primary key (id),
    foreign key (client_properties_id)
            references client_properties(id)
);

CREATE TABLE IF NOT EXISTS client(
    id uuid DEFAULT uuid_generate_v4(),
    client_properties_id uuid not null,
    tag varchar(20),
    time_zone timestamp with time zone not null,
    primary key (id),
    foreign key (client_properties_id)
        references client_properties(id)
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