CREATE TABLE public.users (id serial PRIMARY KEY, first_name VARCHAR ( 50 ), last_name VARCHAR ( 50 ),
email VARCHAR ( 50 ) UNIQUE NOT NULL, password VARCHAR ( 255 ) NOT NULL, created_at timestamp, updated_at timestamp
);