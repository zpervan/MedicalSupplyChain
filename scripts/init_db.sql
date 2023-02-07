-- Import UUID generation functionality while creating a new entry
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_account (
    id uuid DEFAULT uuid_generate_v4 (),
    name text NOT NULL,
    surname text NOT NULL,
    username text UNIQUE NOT NULL,
    password text NOT NULL,
    role text NOT NULL,
    department text NOT NULL,
    created timestamp NOT NULL,
    last_login timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE equipment_inventory (
    id uuid DEFAULT uuid_generate_v4 (),
    quantity int,
    PRIMARY KEY (id)
);

CREATE TABLE equipment_category (
    id uuid DEFAULT uuid_generate_v4 (),
    name text NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE equipment (
    id uuid DEFAULT uuid_generate_v4 (),
    -- TODO: Change equipment category and inventory ID's to be NOT NULL later on
    user_id uuid DEFAULT NULL,
    equipment_inventory_id uuid DEFAULT NULL,
    equipment_category_id uuid DEFAULT NULL,
    name text NOT NULL UNIQUE,
    description text,
    received_date date NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user_account(id),
    FOREIGN KEY (equipment_inventory_id) REFERENCES equipment_inventory(id),
    FOREIGN KEY (equipment_category_id) REFERENCES equipment_category(id)
);

INSERT INTO
    equipment_category (name)
VALUES
    ('Personal protective equipment'),
    ('Injection'),
    ('Wound care'),
    ('Diagnostic'),
    ('Emergency'),
    ('Respiratory care'),
    ('Patient monitoring'),
    ('Medical imaging'),
    ('Laboratory'),
    ('Medical furniture and equipment'),
    ('Medical consumables');