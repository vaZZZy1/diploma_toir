-- Переключаемся на схему execution
SET search_path TO execution;

-- Таблица task_execution
CREATE TABLE task_execution (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    task_id UUID NOT NULL,
    executor_id UUID NOT NULL,
    actual_start_time TIMESTAMP,
    actual_end_time TIMESTAMP,
    status VARCHAR(20) NOT NULL,
    result VARCHAR(20),
    notes TEXT,
    problems_found TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Таблица defect
CREATE TABLE defect (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    aircraft_id UUID NOT NULL,
    detection_date TIMESTAMP NOT NULL,
    detected_by_id UUID NOT NULL,
    category VARCHAR(20) NOT NULL,
    description TEXT NOT NULL,
    location VARCHAR(200),
    component_id UUID,
    status VARCHAR(20) NOT NULL,
    repair_task_id UUID,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Таблица parts_usage
CREATE TABLE parts_usage (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    task_execution_id UUID NOT NULL REFERENCES task_execution(id),
    part_id UUID NOT NULL,
    quantity DECIMAL(10,2) NOT NULL,
    lot_number VARCHAR(50),
    serial_number VARCHAR(50),
    warehouse_id UUID NOT NULL,
    used_by_id UUID NOT NULL,
    used_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Индексы
CREATE INDEX task_execution_task_id_idx ON task_execution(task_id);
CREATE INDEX task_execution_executor_id_idx ON task_execution(executor_id);
CREATE INDEX task_execution_status_idx ON task_execution(status);
CREATE INDEX defect_aircraft_id_idx ON defect(aircraft_id);
CREATE INDEX defect_status_idx ON defect(status);
CREATE INDEX defect_detection_date_idx ON defect(detection_date);
CREATE INDEX parts_usage_task_execution_id_idx ON parts_usage(task_execution_id);
CREATE INDEX parts_usage_part_id_idx ON parts_usage(part_id);
CREATE INDEX parts_usage_used_at_idx ON parts_usage(used_at); 