-- CREATE TABLE IF NOT EXISTS users(
--     id serial,
--     name VARCHAR(255) NOT NULL,
--     telegram_id INTEGER NOT NULL PRIMARY KEY,
--     active boolean DEFAULT 't'
-- );

-- CREATE TABLE IF NOT EXISTS services (
--     id serial PRIMARY KEY,
--     name VARCHAR(255) NOT NULL
-- );

CREATE TABLE IF NOT EXISTS passwords (
    id serial PRIMARY KEY,
    user_id INTEGER NOT NULL, -- REFERENCES users(telegram_id),
    name_of_service VARCHAR(255) NOT NULL, -- REFERENCES services(id),
    username VARCHAR(255),
    pass VARCHAR(255)
);