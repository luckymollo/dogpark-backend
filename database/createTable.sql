CREATE SCHEMA dog_park;

CREATE TABLE dog_owner (
    id serial PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    age INTEGER,
    email VARCHAR (355),
    created_on TIMESTAMP NOT NULL DEFAULT NOW()
);