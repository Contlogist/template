-- создаем таблицу section
CREATE TABLE IF NOT EXISTS section (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    icon VARCHAR(255) NOT NULL
);

-- добавляем записи в таблицу section
INSERT INTO section (name, url, icon)
VALUES
    ('Логистика', 'logistics','logistic.svg'),
    ('Транспорт', 'transport', 'transport.svg'),
    ('Авиация', 'aviation', 'aviation.svg'),
    ('Склад', 'warehouse', 'warehouse.svg');
