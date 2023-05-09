
CREATE TABLE IF NOT EXISTS requests
(
    request_id text default add_request_id() primary key,
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
