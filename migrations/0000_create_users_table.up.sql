CREATE TABLE public.users (id serial PRIMARY KEY, first_name VARCHAR ( 50 ) NOT NULL, last_name VARCHAR ( 50 ) NOT NULL,
email VARCHAR ( 50 ) UNIQUE NOT NULL, password VARCHAR ( 255 ) NOT NULL, deleted BOOLEAN NOT NULL DEFAULT FALSE, created_at timestamp, updated_at timestamp
);