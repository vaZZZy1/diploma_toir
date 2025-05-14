# Схема базы данных системы технического обслуживания воздушных судов

## Общая структура

База данных системы представляет собой реляционную БД PostgreSQL, содержащую несколько взаимосвязанных схем в соответствии с микросервисной архитектурой проекта:

1. `references` - справочные данные и мастер-данные (сервис P1)
2. `planning` - планирование технического обслуживания (сервис P2)
3. `execution` - управление исполнением работ (сервис P3)
4. `inventory` - управление складом и запчастями (сервис P4)
5. `analytics` - аналитика и отчетность (сервис P5)

## Схема `references`

Содержит справочную информацию, используемую всеми сервисами.

### Таблица `aircraft`
Информация о воздушных судах.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор воздушного судна |
| registration_number | VARCHAR(20) | Регистрационный номер ВС |
| aircraft_type_id | UUID | Внешний ключ на тип ВС |
| serial_number | VARCHAR(50) | Серийный номер |
| manufacture_date | DATE | Дата производства |
| operator_id | UUID | Эксплуатант ВС (авиакомпания) |
| base_airport_id | UUID | Базовый аэропорт |
| status | VARCHAR(20) | Статус ВС (active, maintenance, inactive) |
| next_check_date | DATE | Дата следующей проверки |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

### Таблица `aircraft_type`
Типы воздушных судов.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор типа ВС |
| code | VARCHAR(10) | Код типа (например, B737, A320) |
| full_name | VARCHAR(100) | Полное наименование |
| manufacturer_id | UUID | Производитель |
| category | VARCHAR(20) | Категория (passenger, cargo, etc.) |
| max_takeoff_weight | DECIMAL(10,2) | Максимальный взлетный вес |
| max_landing_weight | DECIMAL(10,2) | Максимальный посадочный вес |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

### Таблица `personnel`
Технический персонал.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор сотрудника |
| employee_number | VARCHAR(20) | Табельный номер |
| last_name | VARCHAR(100) | Фамилия |
| first_name | VARCHAR(100) | Имя |
| middle_name | VARCHAR(100) | Отчество |
| position_id | UUID | Должность |
| department_id | UUID | Подразделение |
| qualification_level | VARCHAR(20) | Уровень квалификации |
| email | VARCHAR(100) | Email |
| phone | VARCHAR(20) | Телефон |
| status | VARCHAR(20) | Статус (active, vacation, sick, etc.) |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

### Таблица `maintenance_regulation`
Регламенты технического обслуживания.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор регламента |
| code | VARCHAR(20) | Код регламента |
| name | VARCHAR(200) | Наименование |
| aircraft_type_id | UUID | Тип ВС |
| maintenance_type_id | UUID | Тип ТО (A-check, B-check, C-check, D-check) |
| description | TEXT | Описание |
| periodicity_hours | INT | Периодичность в летных часах |
| periodicity_days | INT | Периодичность в днях |
| periodicity_cycles | INT | Периодичность в циклах |
| estimated_duration_hours | DECIMAL(10,2) | Расчетная продолжительность в часах |
| version | VARCHAR(20) | Версия регламента |
| effective_from | DATE | Дата начала действия |
| effective_to | DATE | Дата окончания действия |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

### Таблица `operator`
Эксплуатанты (авиакомпании).

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| code | VARCHAR(10) | Код авиакомпании |
| name | VARCHAR(200) | Наименование |
| country_id | UUID | Страна регистрации |
| address | TEXT | Юридический адрес |
| contact_person | VARCHAR(200) | Контактное лицо |
| contact_email | VARCHAR(100) | Контактный email |
| contact_phone | VARCHAR(20) | Контактный телефон |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

## Схема `planning`

Содержит данные для планирования технического обслуживания.

### Таблица `maintenance_schedule`
Расписание ТО.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| aircraft_id | UUID | Воздушное судно |
| maintenance_type_id | UUID | Тип ТО |
| regulation_id | UUID | Регламент ТО |
| planned_start_date | TIMESTAMP | Плановая дата начала |
| planned_end_date | TIMESTAMP | Плановая дата окончания |
| status | VARCHAR(20) | Статус (planned, in_progress, completed, canceled) |
| location_id | UUID | Место проведения ТО |
| responsible_person_id | UUID | Ответственное лицо |
| comments | TEXT | Комментарии |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

### Таблица `maintenance_task`
Задачи в рамках ТО.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| schedule_id | UUID | Расписание ТО |
| regulation_task_id | UUID | Задача из регламента |
| planned_start_time | TIMESTAMP | Плановое время начала |
| planned_duration_minutes | INT | Плановая продолжительность в минутах |
| priority | INT | Приоритет |
| status | VARCHAR(20) | Статус (planned, assigned, in_progress, completed, etc.) |
| prerequisites | TEXT | Предварительные условия |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

### Таблица `resource_allocation`
Распределение ресурсов.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| task_id | UUID | Задача ТО |
| resource_type | VARCHAR(20) | Тип ресурса (personnel, equipment, facility) |
| resource_id | UUID | Идентификатор ресурса |
| allocation_start | TIMESTAMP | Время начала выделения |
| allocation_end | TIMESTAMP | Время окончания выделения |
| status | VARCHAR(20) | Статус (planned, confirmed, canceled) |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

## Схема `execution`

Содержит данные о выполнении работ.

### Таблица `task_execution`
Выполнение задач ТО.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| task_id | UUID | Задача ТО |
| executor_id | UUID | Исполнитель |
| actual_start_time | TIMESTAMP | Фактическое время начала |
| actual_end_time | TIMESTAMP | Фактическое время окончания |
| status | VARCHAR(20) | Статус выполнения |
| result | VARCHAR(20) | Результат (success, failure, postponed) |
| notes | TEXT | Примечания |
| problems_found | TEXT | Выявленные проблемы |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

### Таблица `defect`
Дефекты.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| aircraft_id | UUID | Воздушное судно |
| detection_date | TIMESTAMP | Дата обнаружения |
| detected_by_id | UUID | Кем обнаружен |
| category | VARCHAR(20) | Категория дефекта |
| description | TEXT | Описание |
| location | VARCHAR(200) | Место дефекта |
| component_id | UUID | Компонент |
| status | VARCHAR(20) | Статус (open, in_progress, fixed, deferred) |
| repair_task_id | UUID | Задача на устранение |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

### Таблица `parts_usage`
Использование запчастей.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| task_execution_id | UUID | Выполнение задачи |
| part_id | UUID | Запчасть |
| quantity | DECIMAL(10,2) | Количество |
| lot_number | VARCHAR(50) | Номер партии |
| serial_number | VARCHAR(50) | Серийный номер |
| warehouse_id | UUID | Склад |
| used_by_id | UUID | Кем использована |
| used_at | TIMESTAMP | Время использования |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

## Схема `inventory`

Содержит данные о запчастях и управлении складом.

### Таблица `part`
Запчасти.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| part_number | VARCHAR(50) | Номер детали |
| name | VARCHAR(200) | Наименование |
| description | TEXT | Описание |
| category_id | UUID | Категория |
| unit_of_measure | VARCHAR(20) | Единица измерения |
| shelf_life_days | INT | Срок хранения в днях |
| minimum_stock | DECIMAL(10,2) | Минимальный запас |
| reorder_level | DECIMAL(10,2) | Уровень для пополнения |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

### Таблица `inventory_item`
Позиции на складе.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| part_id | UUID | Запчасть |
| warehouse_id | UUID | Склад |
| quantity | DECIMAL(10,2) | Количество |
| lot_number | VARCHAR(50) | Номер партии |
| serial_number | VARCHAR(50) | Серийный номер (для учитываемых поштучно) |
| status | VARCHAR(20) | Статус (available, reserved, quarantine, etc.) |
| location | VARCHAR(50) | Местоположение на складе |
| expiration_date | DATE | Дата истечения срока годности |
| certificate_number | VARCHAR(50) | Номер сертификата |
| purchase_order_id | UUID | Заказ на поставку |
| arrival_date | DATE | Дата поступления |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

### Таблица `warehouse`
Склады.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| code | VARCHAR(20) | Код склада |
| name | VARCHAR(100) | Наименование |
| location_id | UUID | Расположение |
| manager_id | UUID | Руководитель склада |
| is_active | BOOLEAN | Активен/неактивен |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

### Таблица `purchase_order`
Заказы на закупку.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| order_number | VARCHAR(50) | Номер заказа |
| supplier_id | UUID | Поставщик |
| order_date | DATE | Дата заказа |
| expected_delivery_date | DATE | Ожидаемая дата поставки |
| status | VARCHAR(20) | Статус (draft, sent, partially_received, completed) |
| total_amount | DECIMAL(14,2) | Общая сумма |
| currency | VARCHAR(3) | Валюта |
| created_by_id | UUID | Кем создан |
| approved_by_id | UUID | Кем утвержден |
| comments | TEXT | Комментарии |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

## Схема `analytics`

Содержит данные для аналитики и отчетности.

### Таблица `maintenance_metrics`
Метрики ТО.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| aircraft_id | UUID | Воздушное судно |
| date | DATE | Дата |
| scheduled_maintenance_hours | DECIMAL(10,2) | Запланированные часы ТО |
| actual_maintenance_hours | DECIMAL(10,2) | Фактические часы ТО |
| downtime_hours | DECIMAL(10,2) | Часы простоя |
| defects_found | INT | Количество обнаруженных дефектов |
| defects_fixed | INT | Количество устраненных дефектов |
| parts_used | INT | Количество использованных запчастей |
| parts_cost | DECIMAL(14,2) | Стоимость запчастей |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

### Таблица `report`
Отчеты.

| Поле | Тип | Описание |
|------|-----|----------|
| id | UUID | Уникальный идентификатор |
| report_type | VARCHAR(50) | Тип отчета |
| title | VARCHAR(200) | Заголовок |
| params | JSONB | Параметры отчета |
| content | JSONB | Содержимое отчета |
| created_by_id | UUID | Кем создан |
| creation_date | TIMESTAMP | Дата создания |
| period_start | DATE | Начало периода |
| period_end | DATE | Конец периода |
| created_at | TIMESTAMP | Дата создания записи |
| updated_at | TIMESTAMP | Дата обновления записи |

## Индексы и ограничения

Для обеспечения производительности и целостности данных созданы следующие индексы и ограничения:

### Индексы
- По внешним ключам для всех таблиц
- По часто используемым полям для фильтрации (status, date и др.)
- По полям для сортировки (planned_start_date, created_at и др.)

### Ограничения
- Первичные ключи (PRIMARY KEY) для всех таблиц
- Внешние ключи (FOREIGN KEY) со ссылочной целостностью
- Уникальные ограничения (UNIQUE) для кодов, номеров и других уникальных идентификаторов
- Проверочные ограничения (CHECK) для полей с ограниченным набором значений

## Связи между схемами

Каждый микросервис имеет доступ к своей схеме и может обращаться к схеме references для получения справочных данных. Межсервисное взаимодействие осуществляется через API.

## Миграции и версионирование

Для управления схемой базы данных используется система миграций, которая позволяет:
- Отслеживать изменения схемы
- Применять и откатывать изменения
- Управлять версиями схемы

Каждая миграция содержит набор операций по изменению схемы (CREATE, ALTER, DROP) и соответствующие операции отката.