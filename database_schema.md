# Схема базы данных

Упрощенная схема базы данных для микросервисной системы технического обслуживания воздушных судов.

## Схемы и таблицы

### Схема: reference_data
- aircraft: воздушные суда
- aircraft_types: типы ВС
- manufacturers: производители
- airports: аэропорты
- operators: эксплуатанты

### Схема: maintenance_planning
- tasks: задачи обслуживания
- schedules: расписания
- intervals: интервалы

### Схема: work_execution
- work_orders: наряды
- work_items: работы
- inspections: проверки

### Схема: inventory
- parts: запасные части
- inventory: склад
- suppliers: поставщики

### Схема: reporting
- metrics: метрики
- reports: отчеты
- analytics: аналитика