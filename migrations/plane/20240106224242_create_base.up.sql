-- создаем таблицу logistics_task_status
CREATE TABLE IF NOT EXISTS logistics_task_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    icon VARCHAR(255) NOT NULL,
    sort INTEGER NOT NULL,
    status VARCHAR(255) NOT NULL
);

-- добавляем записи в таблицу logistics_task_status
INSERT INTO logistics_task_status (name, color, icon, sort, status)
VALUES
    ('Новая', 'red', 'new.svg', 1, 'active'),
    ('В работе', 'orange', 'in_work.svg', 2, 'active'),
    ('Выполнена', 'yellow', 'done.svg', 3, 'active'),
    ('Отменена', 'green', 'canceled.svg', 4, 'active');

-- создаем таблицу logistics_task
CREATE TABLE IF NOT EXISTS logistics_task (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    date_start DATE NOT NULL,
    date_end DATE NOT NULL,
    status_id INTEGER NOT NULL,
    created_user_id INTEGER NOT NULL,
    responsible_user_id INTEGER NOT NULL,
    FOREIGN KEY (created_user_id) REFERENCES "user"(id),
    FOREIGN KEY (responsible_user_id) REFERENCES "user"(id),
    FOREIGN KEY (status_id) REFERENCES logistics_task_status(id)
);


-- создаем таблицу logistics_request_status
CREATE TABLE IF NOT EXISTS logistics_request_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    icon VARCHAR(255) NOT NULL,
    sort INTEGER NOT NULL,
    status VARCHAR(255) NOT NULL
);

-- добавляем записи в таблицу logistics_request_status
INSERT INTO logistics_request_status (name, color, icon, sort, status)
VALUES
    ('Расчет', 'red', 'calculation.svg', 1, 'active'),
    ('Согласование', 'orange', 'agreement.svg', 2, 'active'),
    ('Распределение', 'yellow', 'distribution.svg', 3, 'active'),
    ('Выполнение', 'green', 'execution.svg', 4, 'active'),
    ('Завершена', 'blue', 'completed.svg', 5, 'active'),
    ('Отменена', 'gray', 'canceled.svg', 6, 'active');

-- создаем таблицу address
CREATE TABLE IF NOT EXISTS address (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    lat VARCHAR(255) NOT NULL,
    lng VARCHAR(255) NOT NULL,
    wait_time INTEGER,
    rating INTEGER,
    company_id INTEGER
);

-- создаем таблицу document
CREATE TABLE IF NOT EXISTS document (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    file VARCHAR(255) NOT NULL,
    date_start DATE NOT NULL,
    date_end DATE NOT NULL,
    created_user_id INTEGER NOT NULL,
    responsible_user_id INTEGER NOT NULL,
    document_type_id INTEGER NOT NULL,
    FOREIGN KEY (created_user_id) REFERENCES "user"(id),
    FOREIGN KEY (responsible_user_id) REFERENCES "user"(id),
    FOREIGN KEY (document_type_id) REFERENCES document_type(id)
    );

-- создаем таблицу cargo
CREATE TABLE IF NOT EXISTS cargo (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    weight INTEGER NOT NULL,
    volume INTEGER NOT NULL,
    count INTEGER NOT NULL,
    company_id INTEGER NOT NULL,
    location_address_id INTEGER NOT NULL,
    destination_address_id INTEGER,
    FOREIGN KEY (company_id) REFERENCES company(id),
    FOREIGN KEY (location_address_id) REFERENCES address(id),
    FOREIGN KEY (destination_address_id) REFERENCES address(id)
);

-- создаем таблицу carrier
CREATE TABLE IF NOT EXISTS carrier (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    site VARCHAR(255) NOT NULL
);

-- создаем таблицу logistics_request
CREATE TABLE IF NOT EXISTS logistics_request (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    date_start DATE NOT NULL,
    date_end DATE NOT NULL,
    status_id INTEGER NOT NULL,
    created_user_id INTEGER,
    company_id INTEGER,
    responsible_user_id INTEGER NOT NULL,
    FOREIGN KEY (created_user_id) REFERENCES "user"(id),
    FOREIGN KEY (responsible_user_id) REFERENCES "user"(id),
    FOREIGN KEY (company_id) REFERENCES company(id),
    FOREIGN KEY (status_id) REFERENCES logistics_request_status(id)
);

-- создаем таблицу logistics_request_cargo
CREATE TABLE IF NOT EXISTS logistics_request_cargo (
    id SERIAL PRIMARY KEY,
    logistics_request_id INTEGER NOT NULL,
    cargo_id INTEGER NOT NULL,
    FOREIGN KEY (logistics_request_id) REFERENCES logistics_request(id),
    FOREIGN KEY (cargo_id) REFERENCES cargo(id)
);

-- создаем таблицу logistics_route_status
CREATE TABLE IF NOT EXISTS logistics_route_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    icon VARCHAR(255) NOT NULL,
    sort INTEGER NOT NULL,
    status VARCHAR(255) NOT NULL
);

-- добавляем записи в таблицу logistics_route_status
INSERT INTO logistics_route_status (name, color, icon, sort, status)
VALUES
    ('Новый', 'red', 'new.svg', 1, 'active'),
    ('В работе', 'orange', 'in_work.svg', 2, 'active'),
    ('Выполнен', 'yellow', 'done.svg', 3, 'active'),
    ('Отменен', 'green', 'canceled.svg', 4, 'active');

-- создаем таблицу logistics_route
CREATE TABLE IF NOT EXISTS logistics_route (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    date_start DATE NOT NULL,
    date_end DATE NOT NULL,
    status_id INTEGER NOT NULL,
    created_user_id INTEGER NOT NULL,
    responsible_user_id INTEGER NOT NULL,
    carrier_id INTEGER NOT NULL,
    FOREIGN KEY (created_user_id) REFERENCES "user"(id),
    FOREIGN KEY (responsible_user_id) REFERENCES "user"(id),
    FOREIGN KEY (carrier_id) REFERENCES carrier(id),
    FOREIGN KEY (status_id) REFERENCES logistics_route_status(id)
);

-- создаем таблицу logistics_route_request
CREATE TABLE IF NOT EXISTS logistics_route_request (
    id SERIAL PRIMARY KEY,
    logistics_route_id INTEGER NOT NULL,
    logistics_request_id INTEGER NOT NULL,
    FOREIGN KEY (logistics_route_id) REFERENCES logistics_route(id),
    FOREIGN KEY (logistics_request_id) REFERENCES logistics_request(id)
);

-- создаем таблицу logistics_route_cargo
CREATE TABLE IF NOT EXISTS logistics_route_cargo (
    id SERIAL PRIMARY KEY,
    logistics_route_id INTEGER NOT NULL,
    cargo_id INTEGER NOT NULL,
    FOREIGN KEY (logistics_route_id) REFERENCES logistics_route(id),
    FOREIGN KEY (cargo_id) REFERENCES cargo(id)
);

-- создаем таблицу logistics_register_status
CREATE TABLE IF NOT EXISTS logistics_register_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    color VARCHAR(255) NOT NULL,
    icon VARCHAR(255) NOT NULL,
    sort INTEGER NOT NULL,
    status VARCHAR(255) NOT NULL
);

-- добавляем записи в таблицу logistics_register_status
INSERT INTO logistics_register_status (name, color, icon, sort, status)
VALUES
    ('Новый', 'red', 'new.svg', 1, 'active'),
    ('В работе', 'orange', 'in_work.svg', 2, 'active'),
    ('Выполнен', 'yellow', 'done.svg', 3, 'active'),
    ('Отменен', 'green', 'canceled.svg', 4, 'active');

-- создаем таблицу logistics_register
CREATE TABLE IF NOT EXISTS logistics_register (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    date_start DATE NOT NULL,
    date_end DATE NOT NULL,
    status_id INTEGER NOT NULL,
    created_user_id INTEGER NOT NULL,
    responsible_user_id INTEGER NOT NULL,
    FOREIGN KEY (created_user_id) REFERENCES "user"(id),
    FOREIGN KEY (responsible_user_id) REFERENCES "user"(id),
    FOREIGN KEY (status_id) REFERENCES logistics_register_status(id)
);

-- создаем таблицу logistics_register_request
CREATE TABLE IF NOT EXISTS logistics_register_request (
    id SERIAL PRIMARY KEY,
    logistics_register_id INTEGER NOT NULL,
    logistics_request_id INTEGER NOT NULL,
    FOREIGN KEY (logistics_register_id) REFERENCES logistics_register(id),
    FOREIGN KEY (logistics_request_id) REFERENCES logistics_request(id)
);

-- создаем таблицу logistics_register_route
CREATE TABLE IF NOT EXISTS logistics_register_route (
    id SERIAL PRIMARY KEY,
    logistics_register_id INTEGER NOT NULL,
    logistics_route_id INTEGER NOT NULL,
    FOREIGN KEY (logistics_register_id) REFERENCES logistics_register(id),
    FOREIGN KEY (logistics_route_id) REFERENCES logistics_route(id)
);

-- создаем таблицу document_type
CREATE TABLE IF NOT EXISTS document_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    icon VARCHAR(255) NOT NULL
);

-- добавляем записи в таблицу document_type
INSERT INTO document_type (name, icon)
VALUES
    ('Договор', 'contract.svg'),
    ('Счет', 'bill.svg'),
    ('Счет-фактура', 'invoice.svg'),
    ('Акт', 'act.svg'),
    ('Справка', 'reference.svg'),
    ('Другое', 'other.svg');


-- создаем таблицу logistics_register_document
CREATE TABLE IF NOT EXISTS logistics_register_document (
    id SERIAL PRIMARY KEY,
    logistics_register_id INTEGER NOT NULL,
    document_id INTEGER NOT NULL,
    FOREIGN KEY (logistics_register_id) REFERENCES logistics_register(id),
    FOREIGN KEY (document_id) REFERENCES document(id)
);





