package converter

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vazy1/reference-service/internal/model"
)

type AircraftConverter struct {
	aircraftTypeCache map[string]uuid.UUID
	manufacturerCache map[string]uuid.UUID
	operatorCache     map[string]uuid.UUID
}

func NewAircraftConverter() *AircraftConverter {
	return &AircraftConverter{
		aircraftTypeCache: make(map[string]uuid.UUID),
		manufacturerCache: make(map[string]uuid.UUID),
		operatorCache:     make(map[string]uuid.UUID),
	}
}

func (c *AircraftConverter) SetAircraftTypeCache(cache map[string]uuid.UUID) {
	c.aircraftTypeCache = cache
}

func (c *AircraftConverter) SetManufacturerCache(cache map[string]uuid.UUID) {
	c.manufacturerCache = cache
}

func (c *AircraftConverter) SetOperatorCache(cache map[string]uuid.UUID) {
	c.operatorCache = cache
}

func (c *AircraftConverter) ToJSON(aircraft model.Aircraft) (string, error) {
	jsonData, err := json.Marshal(aircraft)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации в JSON: %w", err)
	}
	return string(jsonData), nil
}

func (c *AircraftConverter) FromJSON(jsonStr string) (model.Aircraft, error) {
	var aircraft model.Aircraft
	if err := json.Unmarshal([]byte(jsonStr), &aircraft); err != nil {
		return model.Aircraft{}, fmt.Errorf("ошибка десериализации из JSON: %w", err)
	}
	return aircraft, nil
}

func (c *AircraftConverter) ExtractAircraftTypeFromRaw(rawType string) (uuid.UUID, string, error) {
	cleanType := strings.TrimSpace(rawType)
	
	if id, ok := c.aircraftTypeCache[cleanType]; ok {
		return id, cleanType, nil
	}
	
	parts := strings.Split(cleanType, " ")
	if len(parts) > 0 {
		typeCode := parts[0]
		if id, ok := c.aircraftTypeCache[typeCode]; ok {
			return id, typeCode, nil
		}
	}
	
	return uuid.Nil, cleanType, nil
}

func (c *AircraftConverter) ExtractManufacturerFromType(aircraftType string) (uuid.UUID, string, error) {
	manufacturers := map[string]string{
		"A": "Airbus",
		"B": "Boeing",
		"E": "Embraer",
		"C": "Cessna",
		"ATR": "ATR",
		"BE": "Beechcraft",
		"L": "Lockheed Martin",
		"MD": "McDonnell Douglas",
		"S": "Sikorsky",
	}
	
	cleanType := strings.TrimSpace(aircraftType)
	
	for prefix, manufacturer := range manufacturers {
		if strings.HasPrefix(cleanType, prefix) {
			if id, ok := c.manufacturerCache[manufacturer]; ok {
				return id, manufacturer, nil
			}
		}
	}
	
	return uuid.Nil, "", nil
}

func (c *AircraftConverter) EnrichAircraft(aircraft *model.Aircraft) error {
	if aircraft.Type != "" && aircraft.AircraftTypeID == uuid.Nil {
		typeID, _, _ := c.ExtractAircraftTypeFromRaw(aircraft.Type)
		aircraft.AircraftTypeID = typeID
	}
	
	now := time.Now().UTC()
	if aircraft.CreatedAt.IsZero() {
		aircraft.CreatedAt = now
	}
	if aircraft.UpdatedAt.IsZero() {
		aircraft.UpdatedAt = now
	}
	
	if aircraft.ID == uuid.Nil {
		aircraft.ID = uuid.New()
	}
	
	return nil
}