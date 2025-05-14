-- Переключаемся на схему analytics
SET search_path TO analytics;

-- Таблица maintenance_metrics
CREATE TABLE maintenance_metrics (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    aircraft_id UUID NOT NULL,
    date DATE NOT NULL,
    scheduled_maintenance_hours DECIMAL(10,2) NOT NULL,
    actual_maintenance_hours DECIMAL(10,2) NOT NULL,
    downtime_hours DECIMAL(10,2) NOT NULL,
    defects_found INT NOT NULL,
    defects_fixed INT NOT NULL,
    parts_used INT NOT NULL,
    parts_cost DECIMAL(14,2) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT maintenance_metrics_aircraft_date_unique UNIQUE (aircraft_id, date)
);

-- Таблица report
CREATE TABLE report (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    report_type VARCHAR(50) NOT NULL,
    title VARCHAR(200) NOT NULL,
    params JSONB NOT NULL,
    content JSONB NOT NULL,
    created_by_id UUID NOT NULL,
    creation_date TIMESTAMP NOT NULL,
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Индексы
CREATE INDEX maintenance_metrics_aircraft_id_idx ON maintenance_metrics(aircraft_id);
CREATE INDEX maintenance_metrics_date_idx ON maintenance_metrics(date);
CREATE INDEX report_report_type_idx ON report(report_type);
CREATE INDEX report_created_by_id_idx ON report(created_by_id);
CREATE INDEX report_period_idx ON report(period_start, period_end); 