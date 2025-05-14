package model

import (
	"time"

	"github.com/google/uuid"
)

// Aircraft представляет информацию о воздушном судне
type Aircraft struct {
	ID                 uuid.UUID  `json:"id" pg:"id,pk,type:uuid"`
	RegistrationNumber string     `json:"registration_number" pg:"registration_number"`
	Type               string     `json:"type" pg:"type"`
	AircraftTypeID     uuid.UUID  `json:"aircraft_type_id" pg:"aircraft_type_id,type:uuid"`
	SerialNumber       string     `json:"serial_number" pg:"serial_number"`
	ManufactureDate    *time.Time `json:"manufacture_date,omitempty" pg:"manufacture_date"`
	OperatorID         uuid.UUID  `json:"operator_id" pg:"operator_id,type:uuid"`
	BaseAirportID      uuid.UUID  `json:"base_airport_id" pg:"base_airport_id,type:uuid"`
	Status             string     `json:"status" pg:"status"`
	NextCheckDate      *time.Time `json:"next_check_date,omitempty" pg:"next_check_date"`
	CreatedAt          time.Time  `json:"created_at" pg:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at" pg:"updated_at"`
}

// AircraftType представляет тип воздушного судна
type AircraftType struct {
	ID               uuid.UUID `json:"id" pg:"id,pk,type:uuid"`
	Code             string    `json:"code" pg:"code"`
	FullName         string    `json:"full_name" pg:"full_name"`
	ManufacturerID   uuid.UUID `json:"manufacturer_id" pg:"manufacturer_id,type:uuid"`
	Category         string    `json:"category" pg:"category"`
	MaxTakeoffWeight float64   `json:"max_takeoff_weight" pg:"max_takeoff_weight"`
	MaxLandingWeight float64   `json:"max_landing_weight" pg:"max_landing_weight"`
	CreatedAt        time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" pg:"updated_at"`
}

// Manufacturer представляет производителя воздушных судов
type Manufacturer struct {
	ID        uuid.UUID `json:"id" pg:"id,pk,type:uuid"`
	Name      string    `json:"name" pg:"name"`
	Code      string    `json:"code" pg:"code"`
	Country   string    `json:"country" pg:"country"`
	CreatedAt time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt time.Time `json:"updated_at" pg:"updated_at"`
}

// AircraftDataSource представляет источник данных о воздушных судах
type AircraftDataSource struct {
	ID          uuid.UUID `json:"id" pg:"id,pk,type:uuid"`
	Name        string    `json:"name" pg:"name"`
	Type        string    `json:"type" pg:"type"` // csv, xml, api, website
	URL         string    `json:"url" pg:"url"`
	Description string    `json:"description" pg:"description"`
	Active      bool      `json:"active" pg:"active"`
	LastUpdate  time.Time `json:"last_update" pg:"last_update"`
	CreatedAt   time.Time `json:"created_at" pg:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" pg:"updated_at"`
} 