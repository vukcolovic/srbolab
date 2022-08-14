CREATE TABLE public.certificates
(
    id serial PRIMARY KEY,
    source_document_type VARCHAR ( 50 ),
    brand VARCHAR ( 50 ) NOT NULL,
    type_vehicle VARCHAR ( 50 ) NOT NULL,
    variant VARCHAR ( 50 ) NOT NULL,
    version_vehicle VARCHAR ( 50 ) NOT NULL,
    commercial_name VARCHAR ( 50 ),
    estimated_production_year VARCHAR ( 10 ),
    max_mass  VARCHAR ( 10 ),
    running_mass  VARCHAR ( 10 ),
    category VARCHAR ( 30 ),
    bodywork_code VARCHAR ( 80 ),
    axles_tyres_num VARCHAR ( 30 ),
    length  VARCHAR ( 10 ),
    width  VARCHAR ( 10 ),
    height  VARCHAR ( 10 ),
    tyre_wheel VARCHAR ( 30 ),
    engine_code VARCHAR ( 30 ),
    engine_capacity  VARCHAR ( 10 ),
    engine_power  VARCHAR ( 10 ),
    fuel VARCHAR ( 30 ),
    power_weight_ratio VARCHAR ( 30 ),
    seat_number  VARCHAR ( 10 ),
    standing_number  VARCHAR ( 10 ),
    max_speed  VARCHAR ( 10 ),
    gas_level VARCHAR ( 30 ),
    max_laden_mass_axios VARCHAR ( 30 ),
    number_wvta VARCHAR ( 30 ),
    pollution_cert VARCHAR ( 50 ),
    noise_cert VARCHAR ( 50 ),
    coupling_device_approval VARCHAR ( 30 ),
    file_content BYTEA,
    filename VARCHAR ( 50 ),
    created_by INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

ALTER TABLE public.certificates
    ADD CONSTRAINT fk_certificates_created_by FOREIGN KEY (created_by) REFERENCES public.users (id);

ALTER TABLE public.certificates
    ADD UNIQUE (type_vehicle, variant, version_vehicle)

