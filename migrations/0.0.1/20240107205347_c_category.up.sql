-- создаем таблицу category
CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY,
    section_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    icon VARCHAR(255) NOT NULL,
    FOREIGN KEY (section_id) REFERENCES section(id)
);

-- добавляем записи в таблицу category
-- Логистика
INSERT INTO category (section_id, name, url, icon)
VALUES
    (1, 'Задачи', 'tasks', 'tasks.svg'),
    (1, 'Заявки', 'requests', 'requests.svg'),
    (1, 'Маршруты', 'routes', 'routes.svg'),
    (1, 'Реестры', 'registers', 'registers.svg'),
    (1, 'Документы', 'documents', 'documents.svg'),
    (1, 'Справочники', 'directories', 'directories.svg'),
    (1, 'Отчеты', 'reports', 'reports.svg'),
    (1, 'Настройки', 'settings', 'settings.svg'),
-- Транспорт
    (2, 'Задачи', 'tasks', 'tasks.svg'),
    (2, 'Автомобили', 'cars', 'cars.svg'),
    (2, 'Водители', 'drivers', 'drivers.svg'),
    (2, 'Маршруты', 'routes', 'routes.svg'),
    (2, 'Реестры', 'registers', 'registers.svg'),
    (2, 'Документы', 'documents', 'documents.svg'),
-- Авиация
    (3, 'Задачи', 'tasks', 'tasks.svg'),
    (3, 'Самолеты', 'planes', 'planes.svg'),
    (3, 'Пилоты', 'pilots', 'pilots.svg'),
    (3, 'Маршруты', 'routes', 'routes.svg'),
-- Склад
    (4, 'Задачи', 'tasks', 'tasks.svg'),
    (4, 'Склады', 'warehouses', 'warehouses.svg'),
    (4, 'Товары', 'goods', 'goods.svg'),
    (4, 'Реестры', 'registers', 'registers.svg'),
    (4, 'Документы', 'documents', 'documents.svg');