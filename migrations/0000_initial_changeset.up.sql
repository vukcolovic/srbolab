CREATE TABLE public.users
(
    id serial PRIMARY KEY,
    first_name VARCHAR ( 50 ) NOT NULL,
    last_name VARCHAR ( 50 ) NOT NULL,
    email VARCHAR ( 50 ) UNIQUE NOT NULL,
    phone_number VARCHAR ( 30 ),
    contract_number VARCHAR ( 50 ),
    contract_type VARCHAR ( 50 ),
    jmbg VARCHAR ( 30 ),
    adress VARCHAR ( 80 ),
    started_work TIMESTAMP,
    password VARCHAR ( 255 ) NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE public.irregularity_levels
(
    id serial PRIMARY KEY,
    code VARCHAR ( 50 ) NOT NULL,
    created_at TIMESTAMP
);

INSERT INTO public.irregularity_levels (code, created_at) VALUES ('KRITICAN', now());
INSERT INTO public.irregularity_levels (code, created_at) VALUES ('SREDNJI', now());
INSERT INTO public.irregularity_levels (code, created_at) VALUES ('NIZAK', now());

CREATE TABLE public.irregularities
(
    id serial PRIMARY KEY,
    subject VARCHAR ( 255 ) NOT NULL,
    level_id INTEGER NOT NULL,
    controller_id INTEGER,
    created_by INTEGER NOT NULL,
    description VARCHAR ( 1024 ) NOT NULL,
    notice VARCHAR ( 1024 ),
    corrected BOOLEAN NOT NULL DEFAULT FALSE,
    corrected_by INTEGER,
    corrected_date TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

ALTER TABLE public.irregularities
    ADD CONSTRAINT fk_irregularities_levels FOREIGN KEY (level_id) REFERENCES public.irregularity_levels (id);

ALTER TABLE public.irregularities
    ADD CONSTRAINT fk_irregularities_created_by FOREIGN KEY (created_by) REFERENCES public.users (id);

ALTER TABLE public.irregularities
    ADD CONSTRAINT fk_irregularities_corrected_by FOREIGN KEY (corrected_by) REFERENCES public.users (id);

-- CREATE TABLE public.fuel_type
-- (
--     id serial PRIMARY KEY,
--     code VARCHAR ( 50 ) NOT NULL,
--     created_at TIMESTAMP
-- );
--
-- INSERT INTO public.fuel_type (code, created_at) VALUES ('BENZIN', now());
-- INSERT INTO public.fuel_type (code, created_at) VALUES ('DIZEL', now());

CREATE TABLE public.fuel_consumption
(
    id serial PRIMARY KEY,
    date_consumption TIMESTAMP NOT NULL,
    fuel_type VARCHAR ( 50 ) NOT NULL,
    liter NUMERIC(7,2)  NOT NULL,
    price NUMERIC(7,2) ,
    car_registration VARCHAR ( 50 ) NOT NULL,
    poured_by INTEGER NOT NULL,
    bill_file BYTEA,
    created_by INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

ALTER TABLE public.fuel_consumption
    ADD CONSTRAINT fk_fuel_consumption_created_by FOREIGN KEY (created_by) REFERENCES public.users (id);

ALTER TABLE public.fuel_consumption
    ADD CONSTRAINT fk_fuel_consumption_by FOREIGN KEY (poured_by) REFERENCES public.users (id);