CREATE TABLE IF NOT EXISTS ticks.ticks (
	id BIGINT PRIMARY KEY auto_increment NOT NULL,
	`timestamp` BIGINT NULL COMMENT 'Unixtime in milliseconds',
	symbol VARCHAR(100) NULL COMMENT 'name of insrtruments',
	best_bid DOUBLE NULL COMMENT 'the best sell offer',
	best_ask DOUBLE NULL COMMENT 'the best bye offer'
)

CREATE TABLE FIO (
    id SERIAL PRIMARY KEY auto_increment NOT NULL,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    age INT NOT NULL CHECK (age >= 0),
    gender VARCHAR(10) NOT NULL,
    nationality VARCHAR(255) NOT NULL
);
