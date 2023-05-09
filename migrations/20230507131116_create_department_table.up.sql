CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS departments
(
    department_id uuid default uuid_generate_v4() not null primary key,
    name text not null,
    created_time timestamp default now()  not null
);

