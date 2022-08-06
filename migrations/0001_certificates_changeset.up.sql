CREATE TABLE public.certificates
(
    id serial PRIMARY KEY,
    brand VARCHAR ( 50 ) NOT NULL,
    type_vehicle VARCHAR ( 50 ) NOT NULL,
    variant VARCHAR ( 50 ) NOT NULL,
    version_vehicle VARCHAR ( 50 ) NOT NULL,
    commercial_name VARCHAR ( 50 ),
    estimated_production_year INTEGER,
    max_mass INTEGER,
    running_mass INTEGER,
    category VARCHAR ( 30 ),
    bodywork_code VARCHAR ( 80 ),
    axles_tyres_num VARCHAR ( 30 ),
    length INTEGER,
    width INTEGER,
    height INTEGER,
    tyre_wheel VARCHAR ( 30 ),
    engine_code VARCHAR ( 30 ),
    engine_capacity INTEGER,
    engine_power INTEGER,
    fuel VARCHAR ( 30 ),
    power_weight_ratio VARCHAR ( 30 ),
    seat_number INTEGER,
    standing_number INTEGER,
    max_speed INTEGER,
    gas_level VARCHAR ( 30 ),
    max_laden_mass_axios VARCHAR ( 30 ),
    number_wvta VARCHAR ( 30 ),
    pollution_cert VARCHAR ( 30 ),
    noise_cert VARCHAR ( 30 ),
    coupling_device_approval VARCHAR ( 30 ),
    created_by INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

ALTER TABLE public.certificates
    ADD CONSTRAINT fk_certificates_created_by FOREIGN KEY (created_by) REFERENCES public.users (id);

ALTER TABLE public.certificates
    ADD UNIQUE (type_vehicle, variant, version_vehicle)

