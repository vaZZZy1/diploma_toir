package model

import (
	"time"

	"github.com/google/uuid"
)

// MaintenanceSchedule представляет расписание технического обслуживания
type MaintenanceSchedule struct {
	ID                uuid.UUID  `json:"id" pg:"id,pk,type:uuid"`
	AircraftID        uuid.UUID  `json:"aircraft_id" pg:"aircraft_id,type:uuid"`
	MaintenanceTypeID uuid.UUID  `json:"maintenance_type_id" pg:"maintenance_type_id,type:uuid"`
	RegulationID      uuid.UUID  `json:"regulation_id" pg:"regulation_id,type:uuid"`
	PlannedStartDate  time.Time  `json:"planned_start_date" pg:"planned_start_date"`
	PlannedEndDate    time.Time  `json:"planned_end_date" pg:"planned_end_date"`
	Status            string     `json:"status" pg:"status"`
	LocationID        uuid.UUID  `json:"location_id" pg:"location_id,type:uuid"`
	ResponsibleID     uuid.UUID  `json:"responsible_person_id" pg:"responsible_person_id,type:uuid"`
	Comments          string     `json:"comments" pg:"comments"`
	CreatedAt         time.Time  `json:"created_at" pg:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" pg:"updated_at"`
	Tasks             []Task     `json:"tasks" pg:"-"`
	Aircraft          *Aircraft  `json:"aircraft,omitempty" pg:"-"`
}

// Task представляет задачу в рамках технического обслуживания
type Task struct {
	ID                   uuid.UUID `json:"id" pg:"id,pk,type:uuid"`
	ScheduleID           uuid.UUID `json:"schedule_id" pg:"schedule_id,type:uuid"`
	RegulationTaskID     uuid.UUID `json:"regulation_task_id" pg:"regulation_task_id,type:uuid"`
	PlannedStartTime     time.Time `json:"planned_start_time" pg:"planned_start_time"`
	PlannedDurationMins  int       `json:"planned_duration_minutes" pg:"planned_duration_minutes"`
	Priority             int       `json:"priority" pg:"priority"`
	Status               string    `json:"status" pg:"status"`
	Prerequisites        string    `json:"prerequisites" pg:"prerequisites"`
	CreatedAt            time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" pg:"updated_at"`
	ResourceAllocations  []ResourceAllocation `json:"resource_allocations,omitempty" pg:"-"`
}

// ResourceAllocation представляет выделение ресурса для задачи
type ResourceAllocation struct {
	ID             uuid.UUID `json:"id" pg:"id,pk,type:uuid"`
	TaskID         uuid.UUID `json:"task_id" pg:"task_id,type:uuid"`
	ResourceType   string    `json:"resource_type" pg:"resource_type"`
	ResourceID     uuid.UUID `json:"resource_id" pg:"resource_id,type:uuid"`
	AllocationStart time.Time `json:"allocation_start" pg:"allocation_start"`
	AllocationEnd   time.Time `json:"allocation_end" pg:"allocation_end"`
	Status         string    `json:"status" pg:"status"`
	CreatedAt      time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" pg:"updated_at"`
}

// Aircraft представляет краткую информацию о воздушном судне
type Aircraft struct {
	ID                 uuid.UUID `json:"id" pg:"id,pk,type:uuid"`
	RegistrationNumber string    `json:"registration_number" pg:"registration_number"`
	AircraftTypeID     uuid.UUID `json:"aircraft_type_id" pg:"aircraft_type_id,type:uuid"`
	Status             string    `json:"status" pg:"status"`
	NextCheckDate      time.Time `json:"next_check_date" pg:"next_check_date"`
} 