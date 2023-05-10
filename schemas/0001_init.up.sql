CREATE TABLE IF NOT EXISTS users(
    id serial,
    name VARCHAR(255) NOT NULL,
    telegram_id INTEGER NOT NULL PRIMARY KEY,
    active boolean DEFAULT 't'
);

CREATE TABLE IF NOT EXISTS services (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS passwords (
    id serial,
    user_id INTEGER NOT NULL PRIMARY KEY REFERENCES users(telegram_id),
    service_id INTEGER NOT NULL REFERENCES services(id),
    pass VARCHAR(255)
);