package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Report представляет отчет
type Report struct {
	ID           uuid.UUID       `json:"id" pg:"id,pk,type:uuid"`
	ReportType   string          `json:"report_type" pg:"report_type"`
	Title        string          `json:"title" pg:"title"`
	Params       json.RawMessage `json:"params" pg:"params,type:jsonb"`
	Content      json.RawMessage `json:"content" pg:"content,type:jsonb"`
	CreatedByID  uuid.UUID       `json:"created_by_id" pg:"created_by_id,type:uuid"`
	CreationDate time.Time       `json:"creation_date" pg:"creation_date"`
	PeriodStart  time.Time       `json:"period_start" pg:"period_start"`
	PeriodEnd    time.Time       `json:"period_end" pg:"period_end"`
	CreatedAt    time.Time       `json:"created_at" pg:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at" pg:"updated_at"`
}

// ReportParams представляет параметры отчета
type ReportParams struct {
	AircraftID     *uuid.UUID `json:"aircraft_id,omitempty"`
	AircraftTypeID *uuid.UUID `json:"aircraft_type_id,omitempty"`
	OperatorID     *uuid.UUID `json:"operator_id,omitempty"`
	LocationID     *uuid.UUID `json:"location_id,omitempty"`
	Aggregation    string     `json:"aggregation,omitempty"` // daily, weekly, monthly
	IncludeDetails bool       `json:"include_details"`
	SortBy         string     `json:"sort_by,omitempty"`
	SortOrder      string     `json:"sort_order,omitempty"` // asc, desc
	Limit          int        `json:"limit,omitempty"`
}

// MaintenanceMetrics представляет метрики технического обслуживания
type MaintenanceMetrics struct {
	ID                     uuid.UUID `json:"id" pg:"id,pk,type:uuid"`
	AircraftID             uuid.UUID `json:"aircraft_id" pg:"aircraft_id,type:uuid"`
	Date                   time.Time `json:"date" pg:"date"`
	ScheduledMaintenanceHrs float64   `json:"scheduled_maintenance_hours" pg:"scheduled_maintenance_hours"`
	ActualMaintenanceHrs    float64   `json:"actual_maintenance_hours" pg:"actual_maintenance_hours"`
	DowntimeHrs            float64   `json:"downtime_hours" pg:"downtime_hours"`
	DefectsFound           int       `json:"defects_found" pg:"defects_found"`
	DefectsFixed           int       `json:"defects_fixed" pg:"defects_fixed"`
	PartsUsed              int       `json:"parts_used" pg:"parts_used"`
	PartsCost              float64   `json:"parts_cost" pg:"parts_cost"`
	CreatedAt              time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt              time.Time `json:"updated_at" pg:"updated_at"`
}

// MaintenanceEfficiencyReport представляет отчет по эффективности ТО
type MaintenanceEfficiencyReport struct {
	PeriodStart          time.Time                  `json:"period_start"`
	PeriodEnd            time.Time                  `json:"period_end"`
	TotalScheduledHrs    float64                    `json:"total_scheduled_hours"`
	TotalActualHrs       float64                    `json:"total_actual_hours"`
	EfficiencyRatio      float64                    `json:"efficiency_ratio"` // scheduled/actual
	TotalDowntimeHrs     float64                    `json:"total_downtime_hours"`
	AvgTaskCompletionTime float64                   `json:"avg_task_completion_time"`
	DefectsFoundTotal    int                        `json:"defects_found_total"`
	DefectsFixedTotal    int                        `json:"defects_fixed_total"`
	DefectFixRate        float64                    `json:"defect_fix_rate"` // fixed/found
	PartsUsedTotal       int                        `json:"parts_used_total"`
	TotalPartsCost       float64                    `json:"total_parts_cost"`
	AircraftMetrics      map[string]AircraftMetrics `json:"aircraft_metrics,omitempty"`
}

// AircraftMetrics представляет метрики для отдельного воздушного судна
type AircraftMetrics struct {
	AircraftID         uuid.UUID `json:"aircraft_id"`
	RegistrationNumber string    `json:"registration_number"`
	ScheduledHrs       float64   `json:"scheduled_hours"`
	ActualHrs          float64   `json:"actual_hours"`
	DowntimeHrs        float64   `json:"downtime_hours"`
	DefectsFound       int       `json:"defects_found"`
	DefectsFixed       int       `json:"defects_fixed"`
	PartsUsed          int       `json:"parts_used"`
	PartsCost          float64   `json:"parts_cost"`
	EfficiencyRatio    float64   `json:"efficiency_ratio"`
} 