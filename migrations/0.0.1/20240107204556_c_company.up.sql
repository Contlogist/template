-- создаем таблицу company
CREATE TABLE IF NOT EXISTS company (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    site VARCHAR(255) NOT NULL
);

-- создаем компанию ООО Континент
INSERT INTO company (name, address, phone, email, site)
VALUES
    ('ООО Континент', 'г. Москва, ул. Ленина, д. 1', '79261234567', 'r.abramov@contlogist.ru', 'https://contlogist.ru/');

-- создаем таблицу company_level
CREATE TABLE IF NOT EXISTS company_level (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    level INTEGER NOT NULL
);

-- добавляем записи в таблицу company_level
INSERT INTO company_level (name, level)
VALUES
    ('Клиент', 0),
    ('Партнер', 1);

-- создаем таблицу company_access
CREATE TABLE IF NOT EXISTS company_access (
    id SERIAL PRIMARY KEY,
    company_id INTEGER NOT NULL,
    company_level_id INTEGER NOT NULL,
    section_id INTEGER NOT NULL,
    FOREIGN KEY (company_id) REFERENCES company(id),
    FOREIGN KEY (section_id) REFERENCES section(id),
    FOREIGN KEY (company_level_id) REFERENCES company_level(id)
);

-- добавляем запись в таблицу company_access
INSERT INTO company_access (company_id, section_id, company_level_id)
VALUES
    (1, 1, 1),
    (1, 2, 1),
    (1, 3, 1),
    (1, 4, 1);