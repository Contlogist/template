-- создаем таблицу user_level
CREATE TABLE IF NOT EXISTS user_level (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- создаем таблицу user
CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    company_id INTEGER NOT NULL,
    refresh_token VARCHAR(255),
    FOREIGN KEY (company_id) REFERENCES company(id)
);

-- добавляем записи в таблицу user_level
INSERT INTO user_level (name)
VALUES
    ('Администратор'),
    ('Менеджер'),
    ('Клиент');

-- создаем таблицу user_access
CREATE TABLE IF NOT EXISTS user_access (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    user_level_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES "user"(id),
    FOREIGN KEY (category_id) REFERENCES category(id),
    FOREIGN KEY (user_level_id) REFERENCES user_level(id)
);


-- добавляем запись в таблицу user
INSERT INTO "user" (name, email, password, company_id)
VALUES
    ('Роман Абрамов', 'r.abramov@contlogist.ru', '123456', 1);

-- добавляем записи в таблицу user_access
-- Логистика
INSERT INTO user_access (user_id, category_id, user_level_id)
VALUES
    (1, 1, 1),
    (1, 2, 1),
    (1, 3, 1),
    (1, 4, 1),
    (1, 5, 1),
    (1, 6, 1),
    (1, 7, 1),
    (1, 8, 1),
-- Транспорт
    (1, 9, 1),
    (1, 10, 1),
    (1, 11, 1),
    (1, 12, 1),
    (1, 13, 1),
    (1, 14, 1),
-- Авиация
    (1, 15, 1),
    (1, 16, 1),
    (1, 17, 1),
    (1, 18, 1),
-- Склад
    (1, 19, 1),
    (1, 20, 1),
    (1, 21, 1),
    (1, 22, 1),
    (1, 23, 1);