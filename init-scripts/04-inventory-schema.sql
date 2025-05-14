-- Переключаемся на схему inventory
SET search_path TO inventory;

-- Таблица part
CREATE TABLE part (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    part_number VARCHAR(50) NOT NULL,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    category_id UUID,
    unit_of_measure VARCHAR(20) NOT NULL,
    shelf_life_days INT,
    minimum_stock DECIMAL(10,2),
    reorder_level DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT part_part_number_unique UNIQUE (part_number)
);

-- Таблица inventory_item
CREATE TABLE inventory_item (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    part_id UUID NOT NULL REFERENCES part(id),
    warehouse_id UUID NOT NULL,
    quantity DECIMAL(10,2) NOT NULL,
    lot_number VARCHAR(50),
    serial_number VARCHAR(50),
    status VARCHAR(20) NOT NULL,
    location VARCHAR(50),
    expiration_date DATE,
    certificate_number VARCHAR(50),
    purchase_order_id UUID,
    arrival_date DATE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Таблица warehouse
CREATE TABLE warehouse (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(20) NOT NULL,
    name VARCHAR(100) NOT NULL,
    location_id UUID,
    manager_id UUID,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT warehouse_code_unique UNIQUE (code)
);

-- Таблица purchase_order
CREATE TABLE purchase_order (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    order_number VARCHAR(50) NOT NULL,
    supplier_id UUID NOT NULL,
    order_date DATE NOT NULL,
    expected_delivery_date DATE,
    status VARCHAR(20) NOT NULL,
    total_amount DECIMAL(14,2),
    currency VARCHAR(3),
    created_by_id UUID NOT NULL,
    approved_by_id UUID,
    comments TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT purchase_order_number_unique UNIQUE (order_number)
);

-- Индексы
CREATE INDEX part_category_id_idx ON part(category_id);
CREATE INDEX inventory_item_part_id_idx ON inventory_item(part_id);
CREATE INDEX inventory_item_warehouse_id_idx ON inventory_item(warehouse_id);
CREATE INDEX inventory_item_status_idx ON inventory_item(status);
CREATE INDEX inventory_item_expiration_date_idx ON inventory_item(expiration_date);
CREATE INDEX warehouse_is_active_idx ON warehouse(is_active);
CREATE INDEX purchase_order_supplier_id_idx ON purchase_order(supplier_id);
CREATE INDEX purchase_order_status_idx ON purchase_order(status);
CREATE INDEX purchase_order_date_idx ON purchase_order(order_date); 