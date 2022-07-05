CREATE TABLE public.users
(
    id serial PRIMARY KEY,
    first_name VARCHAR ( 50 ) NOT NULL,
    last_name VARCHAR ( 50 ) NOT NULL,
    email VARCHAR ( 50 ) UNIQUE NOT NULL,
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