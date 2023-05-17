CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS requests
(
    request_id  uuid default uuid_generate_v4() not null primary key,
    request_code text default add_request_code() not null,
    employee_id uuid not null,
    department_id uuid not null,
    equipment_id uuid not null,
    quanity integer not null,
    price integer not null,
    created_time timestamp default now() not null
);

-- create index requests_request_id_index
--     on requests USING hash(request_id);

-- create index requests_employee_id_index
--     on requests USING hash(employee_id);

-- create index requests_equipment_id_index
--     on requests USING hash(equipment_id);

-- create index requests_department_id_index
--     on requests USING hash(department_id);
