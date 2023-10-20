CREATE TABLE IF NOT EXISTS fio (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    age INT NOT NULL CHECK (age >= 0),
    gender VARCHAR(10) NOT NULL,
    nationality VARCHAR(255) NOT NULL
);
