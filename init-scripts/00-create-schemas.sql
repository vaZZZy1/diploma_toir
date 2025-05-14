-- Создание схем для каждого микросервиса
CREATE SCHEMA IF NOT EXISTS references;
CREATE SCHEMA IF NOT EXISTS planning;
CREATE SCHEMA IF NOT EXISTS execution;
CREATE SCHEMA IF NOT EXISTS inventory;
CREATE SCHEMA IF NOT EXISTS analytics;

-- Создание расширения для генерации UUID
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Установка пути поиска схем
SET search_path TO references, planning, execution, inventory, analytics, public;

-- Предоставление прав на использование схем
GRANT USAGE ON SCHEMA references TO postgres;
GRANT USAGE ON SCHEMA planning TO postgres;
GRANT USAGE ON SCHEMA execution TO postgres;
GRANT USAGE ON SCHEMA inventory TO postgres;
GRANT USAGE ON SCHEMA analytics TO postgres; 