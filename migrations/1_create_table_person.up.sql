CREATE TABLE person (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    age INT,
    birth_date DATE,
    address VARCHAR(255)
);