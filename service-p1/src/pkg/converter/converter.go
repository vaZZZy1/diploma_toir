package converter

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vazy1/reference-service/internal/model"
)

// AircraftConverter предоставляет методы для конвертации данных воздушных судов
type AircraftConverter struct {
	// Кеши для поиска идентификаторов
	aircraftTypeCache map[string]uuid.UUID
	manufacturerCache map[string]uuid.UUID
	operatorCache     map[string]uuid.UUID
}

// NewAircraftConverter создает новый экземпляр конвертера
func NewAircraftConverter() *AircraftConverter {
	return &AircraftConverter{
		aircraftTypeCache: make(map[string]uuid.UUID),
		manufacturerCache: make(map[string]uuid.UUID),
		operatorCache:     make(map[string]uuid.UUID),
	}
}

// SetAircraftTypeCache устанавливает кеш для поиска типов ВС
func (c *AircraftConverter) SetAircraftTypeCache(cache map[string]uuid.UUID) {
	c.aircraftTypeCache = cache
}

// SetManufacturerCache устанавливает кеш для поиска производителей
func (c *AircraftConverter) SetManufacturerCache(cache map[string]uuid.UUID) {
	c.manufacturerCache = cache
}

// SetOperatorCache устанавливает кеш для поиска операторов
func (c *AircraftConverter) SetOperatorCache(cache map[string]uuid.UUID) {
	c.operatorCache = cache
}

// ToJSON конвертирует Aircraft в JSON строку
func (c *AircraftConverter) ToJSON(aircraft model.Aircraft) (string, error) {
	jsonData, err := json.Marshal(aircraft)
	if err != nil {
		return "", fmt.Errorf("ошибка сериализации в JSON: %w", err)
	}
	return string(jsonData), nil
}

// FromJSON конвертирует JSON строку в Aircraft
func (c *AircraftConverter) FromJSON(jsonStr string) (model.Aircraft, error) {
	var aircraft model.Aircraft
	if err := json.Unmarshal([]byte(jsonStr), &aircraft); err != nil {
		return model.Aircraft{}, fmt.Errorf("ошибка десериализации из JSON: %w", err)
	}
	return aircraft, nil
}

// ExtractAircraftTypeFromRaw извлекает тип ВС из строки и пытается найти его в кеше
func (c *AircraftConverter) ExtractAircraftTypeFromRaw(rawType string) (uuid.UUID, string, error) {
	// Очищаем и нормализуем строку
	cleanType := strings.TrimSpace(rawType)
	
	// Проверяем наличие в кеше
	if id, ok := c.aircraftTypeCache[cleanType]; ok {
		return id, cleanType, nil
	}
	
	// Пытаемся извлечь код типа
	parts := strings.Split(cleanType, " ")
	if len(parts) > 0 {
		typeCode := parts[0]
		if id, ok := c.aircraftTypeCache[typeCode]; ok {
			return id, typeCode, nil
		}
	}
	
	// Если не нашли, возвращаем нулевой ID и оригинальную строку
	return uuid.Nil, cleanType, nil
}

// ExtractManufacturerFromType пытается извлечь производителя из типа ВС
func (c *AircraftConverter) ExtractManufacturerFromType(aircraftType string) (uuid.UUID, string, error) {
	// Основные производители
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
	
	// Проверяем префиксы типа ВС
	for prefix, manufacturer := range manufacturers {
		if strings.HasPrefix(cleanType, prefix) {
			if id, ok := c.manufacturerCache[manufacturer]; ok {
				return id, manufacturer, nil
			}
		}
	}
	
	return uuid.Nil, "", nil
}

// EnrichAircraft обогащает данные воздушного судна дополнительной информацией
func (c *AircraftConverter) EnrichAircraft(aircraft *model.Aircraft) error {
	// Если тип есть, но ID типа нет, пытаемся найти ID
	if aircraft.Type != "" && aircraft.AircraftTypeID == uuid.Nil {
		typeID, _, _ := c.ExtractAircraftTypeFromRaw(aircraft.Type)
		aircraft.AircraftTypeID = typeID
	}
	
	// Устанавливаем текущее время для полей Created/Updated если не установлены
	now := time.Now().UTC()
	if aircraft.CreatedAt.IsZero() {
		aircraft.CreatedAt = now
	}
	if aircraft.UpdatedAt.IsZero() {
		aircraft.UpdatedAt = now
	}
	
	// Генерируем ID если отсутствует
	if aircraft.ID == uuid.Nil {
		aircraft.ID = uuid.New()
	}
	
	return nil
} 