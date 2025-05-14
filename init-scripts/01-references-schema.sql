-- Переключаемся на схему references
SET search_path TO references;

-- Таблица aircraft_type
CREATE TABLE aircraft_type (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(10) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    manufacturer_id UUID,
    category VARCHAR(20),
    max_takeoff_weight DECIMAL(10,2),
    max_landing_weight DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT aircraft_type_code_unique UNIQUE (code)
);

-- Таблица aircraft
CREATE TABLE aircraft (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    registration_number VARCHAR(20) NOT NULL,
    aircraft_type_id UUID NOT NULL REFERENCES aircraft_type(id),
    serial_number VARCHAR(50) NOT NULL,
    manufacture_date DATE,
    operator_id UUID,
    base_airport_id UUID,
    status VARCHAR(20) NOT NULL,
    next_check_date DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT aircraft_registration_unique UNIQUE (registration_number)
);

-- Таблица personnel
CREATE TABLE personnel (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_number VARCHAR(20) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    position_id UUID,
    department_id UUID,
    qualification_level VARCHAR(20),
    email VARCHAR(100),
    phone VARCHAR(20),
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT personnel_employee_number_unique UNIQUE (employee_number)
);

-- Таблица maintenance_regulation
CREATE TABLE maintenance_regulation (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(20) NOT NULL,
    name VARCHAR(200) NOT NULL,
    aircraft_type_id UUID NOT NULL REFERENCES aircraft_type(id),
    maintenance_type_id UUID,
    description TEXT,
    periodicity_hours INT,
    periodicity_days INT,
    periodicity_cycles INT,
    estimated_duration_hours DECIMAL(10,2),
    version VARCHAR(20),
    effective_from DATE,
    effective_to DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT maintenance_regulation_code_version_unique UNIQUE (code, version)
);

-- Таблица operator
CREATE TABLE operator (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(10) NOT NULL,
    name VARCHAR(200) NOT NULL,
    country_id UUID,
    address TEXT,
    contact_person VARCHAR(200),
    contact_email VARCHAR(100),
    contact_phone VARCHAR(20),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT operator_code_unique UNIQUE (code)
);

-- Индексы
CREATE INDEX aircraft_type_id_idx ON aircraft(aircraft_type_id);
CREATE INDEX aircraft_operator_id_idx ON aircraft(operator_id);
CREATE INDEX aircraft_status_idx ON aircraft(status);
CREATE INDEX personnel_status_idx ON personnel(status);
CREATE INDEX maintenance_regulation_aircraft_type_id_idx ON maintenance_regulation(aircraft_type_id);
CREATE INDEX maintenance_regulation_effective_dates_idx ON maintenance_regulation(effective_from, effective_to); 