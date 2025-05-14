-- Переключаемся на схему planning
SET search_path TO planning;

-- Таблица maintenance_schedule
CREATE TABLE maintenance_schedule (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    aircraft_id UUID NOT NULL,
    maintenance_type_id UUID NOT NULL,
    regulation_id UUID NOT NULL,
    planned_start_date TIMESTAMP NOT NULL,
    planned_end_date TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL,
    location_id UUID,
    responsible_person_id UUID,
    comments TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Таблица maintenance_task
CREATE TABLE maintenance_task (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    schedule_id UUID NOT NULL REFERENCES maintenance_schedule(id),
    regulation_task_id UUID,
    planned_start_time TIMESTAMP NOT NULL,
    planned_duration_minutes INT NOT NULL,
    priority INT DEFAULT 0,
    status VARCHAR(20) NOT NULL,
    prerequisites TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Таблица resource_allocation
CREATE TABLE resource_allocation (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    task_id UUID NOT NULL REFERENCES maintenance_task(id),
    resource_type VARCHAR(20) NOT NULL,
    resource_id UUID NOT NULL,
    allocation_start TIMESTAMP NOT NULL,
    allocation_end TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Индексы
CREATE INDEX maintenance_schedule_aircraft_id_idx ON maintenance_schedule(aircraft_id);
CREATE INDEX maintenance_schedule_status_idx ON maintenance_schedule(status);
CREATE INDEX maintenance_schedule_date_idx ON maintenance_schedule(planned_start_date);
CREATE INDEX maintenance_task_schedule_id_idx ON maintenance_task(schedule_id);
CREATE INDEX maintenance_task_status_idx ON maintenance_task(status);
CREATE INDEX resource_allocation_task_id_idx ON resource_allocation(task_id);
CREATE INDEX resource_allocation_resource_id_idx ON resource_allocation(resource_id);
CREATE INDEX resource_allocation_date_idx ON resource_allocation(allocation_start, allocation_end); 